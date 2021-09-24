package redisCache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"runtime"
	"time"
)

type redisLock struct {
	redis redis.Cmdable
}

type LockType int

const (
	Writer LockType = 1
	Read   LockType = 2

	DefaultLockTtl   = time.Minute * 10
	DefaultLockValue = "0"
)

type RedisLock interface {
	// 加锁, 返回加锁是否成功
	TryLock(ctx context.Context, key string, val interface{}, ttl time.Duration) (bool, []byte, error)
	// 解锁
	UnLock(ctx context.Context, key string, val interface{}) error
	// 等待释放锁
	WaitUnLock(ctx context.Context, key string) (err error)
	// 续约
	RenewLock(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error)
	// 加锁, 如果有锁，等待锁释放再添加锁
	WaitLock(ctx context.Context, key string, val interface{}, ttl time.Duration) (bool, error)
}

func NewRedisLock(redis redis.Cmdable) RedisLock {
	return &redisLock{
		redis: redis,
	}
}

// TryLock 加锁, 返回加锁是否成功
func (c redisLock) TryLock(ctx context.Context, key string, val interface{}, ttl time.Duration) (bool, []byte, error) {
	var lockKey = NewLockKey(key)
	var result []byte
	success, err := c.redis.SetNX(ctx, lockKey, val, ttl).Result()
	if err != nil {
		return false, nil, err
	}
	if !success {
		result, err = c.redis.Get(ctx, lockKey).Bytes()
		return success, result, err
	}
	return success, result, err
}

// WaitLock 加锁, 如果有锁，等待锁释放再添加锁
func (c redisLock) WaitLock(ctx context.Context, key string, val interface{}, ttl time.Duration) (bool, error) {
	var lockKey = NewLockKey(key)
	err := c.WaitUnLock(ctx, lockKey)
	if err != nil {
		return false, err
	}
	lua := redis.NewScript("if redis.call('exists', KEYS[1]) == 0 " +
		"then return redis.call('set', KEYS[1], ARGV[1], 'EX', ARGV[2]) " +
		"else return 'NO' end")
	cmd := lua.Run(ctx, c.redis, []string{lockKey}, val, int(ttl/time.Second))
	if cmd.Err() != nil {
		return false, cmd.Err()
	} else {
		result := cmd.String()
		if result == "OK" {
			return true, cmd.Err()
		} else {
			return false, cmd.Err()
		}
	}
}

// UnLock 解锁
func (c redisLock) UnLock(ctx context.Context, key string, val interface{}) error {
	var lockKey = NewLockKey(key)
	luaDel := redis.NewScript("if redis.call('get',KEYS[1]) == ARGV[1] then " +
		"return redis.call('del',KEYS[1]) else return 0 end")
	return luaDel.Run(ctx, c.redis, []string{lockKey}, val).Err()
}

// WaitUnLock 等待释放锁
func (c redisLock) WaitUnLock(ctx context.Context, key string) (err error) {
	var lockKey = NewLockKey(key)
	ttl, err := c.redis.TTL(ctx, lockKey).Result()
	if err != nil {
		return err
	}
	if ttl > 0 {
		time.Sleep(ttl)
		runtime.Gosched()
		return c.WaitUnLock(ctx, key)
	}
	return nil
}

// RenewLock 续约
func (c redisLock) RenewLock(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error) {
	return c.redis.Set(ctx, NewLockKey(key), value, ttl).Err()
}

// NewLockKey 拼接锁key
func NewLockKey(key string) (lockKey string) {
	return key + "_lock"
}
