package redisCache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/bufpool"
	"golang.org/x/sync/singleflight"
	"log"
	"sync/atomic"
	"time"
)

type Item struct {
	Ctx context.Context

	Key   string
	Value interface{}

	// TTL is the cache expiration time.
	// Default TTL is 1 hour.
	TTL time.Duration

	// Do returns value to be cached.
	Do func(*Item) (interface{}, error)

	// SetXX only sets the key if it already exists.
	SetXX bool

	// SetNX only sets the key if it does not already exist.
	SetNX bool

	// SkipLocalCache skips local cache as if it is not set.
	SkipLocalCache bool
}

func (item *Item) Context() context.Context {
	if item.Ctx == nil {
		return context.Background()
	}
	return item.Ctx
}

func (item *Item) value() (interface{}, error) {
	if item.Do != nil {
		return item.Do(item)
	}
	if item.Value != nil {
		return item.Value, nil
	}
	return nil, nil
}

func (item *Item) ttl() time.Duration {
	const defaultTTL = time.Hour

	if item.TTL < 0 {
		return 0
	}

	if item.TTL != 0 {
		if item.TTL < time.Second {
			log.Printf("too short TTL for key=%q: %s", item.Key, item.TTL)
			return defaultTTL
		}
		return item.TTL
	}

	return defaultTTL
}

var (
	//ErrCacheMiss          = errors.New("cache: key is missing")
	errRedisLocalCacheNil = errors.New("cache: both Redis and LocalCache are nil")
)

type MarshalFunc func(interface{}) ([]byte, error)
type UnmarshalFunc func([]byte, interface{}) error

type Options struct {
	Redis        redis.Cmdable
	LocalCache   LocalCache
	StatsEnabled bool
	Marshal      MarshalFunc
	Unmarshal    UnmarshalFunc
}

type Cache struct {
	opt *Options

	group   singleflight.Group
	bufpool bufpool.Pool

	marshal   MarshalFunc
	unmarshal UnmarshalFunc

	hits   uint64
	misses uint64
}

func New(opt *Options) *Cache {
	cacher := &Cache{
		opt: opt,
	}

	if opt.Marshal == nil {
		cacher.marshal = Marshal
	} else {
		cacher.marshal = opt.Marshal
	}

	if opt.Unmarshal == nil {
		cacher.unmarshal = Unmarshal
	} else {
		cacher.unmarshal = opt.Unmarshal
	}

	return cacher
}

var (
	ErrCacheMiss = errors.New("cache: key is missing")
)

// Set caches the item.
func (cd *Cache) Set(ctx context.Context, item *Item) error {
	var err error
	if err != nil {
		return err
	}
	_, _, err = cd.set(item)
	return err
}

func (cd *Cache) set(item *Item) ([]byte, bool, error) {
	value, err := item.value()
	if err != nil {
		return nil, false, err
	}

	b, err := cd.Marshal(value)
	if err != nil {
		return nil, false, err
	}

	if cd.opt.LocalCache != nil && !item.SkipLocalCache {
		cd.opt.LocalCache.Set(item.Key, b)
	}

	if cd.opt.Redis == nil {
		if cd.opt.LocalCache == nil {
			return b, true, errRedisLocalCacheNil
		}
		return b, true, nil
	}

	ttl := item.ttl()
	if ttl == 0 {
		return b, true, nil
	}

	if item.SetXX {
		return b, true, cd.opt.Redis.SetXX(item.Ctx, item.Key, b, ttl).Err()
	}
	if item.SetNX {
		return b, true, cd.opt.Redis.SetNX(item.Ctx, item.Key, b, ttl).Err()
	}
	return b, true, cd.opt.Redis.Set(item.Ctx, item.Key, b, ttl).Err()
}

// Exists reports whether value for the given key exists.
func (cd *Cache) Exists(ctx context.Context, key string) bool {
	result, _ := cd.opt.Redis.Exists(ctx, key).Result()
	return result == 1
}

// Get gets the value for the given key.
func (cd *Cache) Get(ctx context.Context, key string, value interface{}) error {
	return cd.get(ctx, key, value, false)
}

// GetSkippingLocalCache gets the value for the given key skipping local cache.
func (cd *Cache) GetSkippingLocalCache(
	ctx context.Context, key string, value interface{},
) error {
	return cd.get(ctx, key, value, true)
}

func (cd *Cache) get(
	ctx context.Context,
	key string,
	value interface{},
	skipLocalCache bool,
) error {
	b, err := cd.getBytes(ctx, key, skipLocalCache)
	if err != nil {
		return err
	}
	return cd.unmarshal(b, value)
}

func (cd *Cache) getBytes(ctx context.Context, key string, skipLocalCache bool) ([]byte, error) {
	if !skipLocalCache && cd.opt.LocalCache != nil {
		b, ok := cd.opt.LocalCache.Get(key)
		if ok {
			return b, nil
		}
	}

	if cd.opt.Redis == nil {
		if cd.opt.LocalCache == nil {
			return nil, errRedisLocalCacheNil
		}
		return nil, ErrCacheMiss
	}

	b, err := cd.opt.Redis.Get(ctx, key).Bytes()
	if err != nil {
		if cd.opt.StatsEnabled {
			atomic.AddUint64(&cd.misses, 1)
		}
		return nil, err
	}

	if cd.opt.StatsEnabled {
		atomic.AddUint64(&cd.hits, 1)
	}

	if !skipLocalCache && cd.opt.LocalCache != nil {
		cd.opt.LocalCache.Set(key, b)
	}
	return b, nil
}

// Once gets the item.Value for the given item.Key from the cache or
// executes, caches, and returns the results of the given item.Func,
// making sure that only one execution is in-flight for a given item.Key
// at a time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (cd *Cache) Once(item *Item) error {
	b, cached, err := cd.getSetItemBytesOnce(item)
	if err != nil {
		return err
	}

	if item.Value == nil || len(b) == 0 {
		return nil
	}

	if err := cd.unmarshal(b, item.Value); err != nil {
		if cached {
			_ = cd.Delete(item.Context(), item.Key)
			return cd.Once(item)
		}
		return err
	}

	return nil
}

func (cd *Cache) getSetItemBytesOnce(item *Item) (b []byte, cached bool, err error) {
	if cd.opt.LocalCache != nil {
		b, ok := cd.opt.LocalCache.Get(item.Key)
		if ok {
			return b, true, nil
		}
	}

	v, err, _ := cd.group.Do(item.Key, func() (interface{}, error) {
		b, err := cd.getBytes(item.Context(), item.Key, item.SkipLocalCache)
		if err == nil {
			cached = true
			return b, nil
		}

		b, ok, err := cd.set(item)
		if ok {
			return b, nil
		}
		return nil, err
	})
	if err != nil {
		return nil, false, err
	}
	return v.([]byte), cached, nil
}

func (cd *Cache) Delete(ctx context.Context, key string) error {
	if cd.opt.LocalCache != nil {
		cd.opt.LocalCache.Del(key)
	}

	if cd.opt.Redis == nil {
		if cd.opt.LocalCache == nil {
			return errRedisLocalCacheNil
		}
		return nil
	}

	_, err := cd.opt.Redis.Del(ctx, key).Result()
	return err
}

func (cd *Cache) DeleteFromLocalCache(key string) {
	if cd.opt.LocalCache != nil {
		cd.opt.LocalCache.Del(key)
	}
}

func (cd *Cache) Marshal(value interface{}) ([]byte, error) {
	return cd.marshal(value)
}

func (cd *Cache) Unmarshal(b []byte, value interface{}) error {
	return cd.unmarshal(b, value)
}

func (cd *Cache) SetS(ctx context.Context, items []*Item) error {
	var err error
	for _, item := range items {
		err = cd.Set(ctx, item)
		if err != nil {
			return err
		}
	}
	return err
}

func (cd *Cache) setByTx(item *Item, pipeliner redis.Pipeliner) ([]byte, bool, error) {
	value, err := item.value()
	if err != nil {
		return nil, false, err
	}

	b, err := cd.Marshal(value)
	if err != nil {
		return nil, false, err
	}

	if cd.opt.LocalCache != nil && !item.SkipLocalCache {
		cd.opt.LocalCache.Set(item.Key, b)
	}

	if pipeliner == nil {
		if cd.opt.LocalCache == nil {
			return b, true, errRedisLocalCacheNil
		}
		return b, true, nil
	}

	ttl := item.ttl()
	if ttl == 0 {
		return b, true, nil
	}

	if item.SetXX {
		return b, true, pipeliner.SetXX(item.Ctx, item.Key, b, ttl).Err()
	}
	if item.SetNX {
		return b, true, pipeliner.SetNX(item.Ctx, item.Key, b, ttl).Err()
	}
	return b, true, pipeliner.Set(item.Ctx, item.Key, b, ttl).Err()
}

func (cd *Cache) GetS(ctx context.Context, keys []string, skipLocalCache bool) ([][]byte, error) {
	var err error
	var tx = cd.opt.Redis.TxPipeline()
	var caches = make([][]byte, 0)
	for _, item := range keys {
		cache, err := cd.getBytesByTx(ctx, item, skipLocalCache, tx)
		if err != nil {
			tx.Discard()
			return nil, err
		}
		if len(cache) != 0 { // 如果有本地缓存，先添加
			caches = append(caches, cache)
		}
	}
	result, err := tx.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	} else if errors.Is(err, redis.Nil) { // 过滤没有查询到的错误
		err = nil
	}
	for _, item := range result {
		res := item.(*redis.StringCmd)
		info, err := res.Bytes()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, err
		}
		if errors.Is(err, redis.Nil) || len(info) == 0 {
			continue
		}
		caches = append(caches, info)
	}
	return caches, err
}

func (cd *Cache) getBytesByTx(ctx context.Context, key string, skipLocalCache bool, pipeliner redis.Pipeliner) ([]byte, error) {
	if !skipLocalCache && cd.opt.LocalCache != nil {
		b, ok := cd.opt.LocalCache.Get(key)
		if ok {
			return b, nil
		}
	}

	if pipeliner == nil {
		if cd.opt.LocalCache == nil {
			return nil, errRedisLocalCacheNil
		}
		return nil, ErrCacheMiss
	}

	b, err := pipeliner.Get(ctx, key).Bytes()
	if err != nil {
		if cd.opt.StatsEnabled {
			atomic.AddUint64(&cd.misses, 1)
		}
		return nil, err
	}

	if cd.opt.StatsEnabled {
		atomic.AddUint64(&cd.hits, 1)
	}

	if !skipLocalCache && cd.opt.LocalCache != nil {
		cd.opt.LocalCache.Set(key, b)
	}
	return b, nil
}
