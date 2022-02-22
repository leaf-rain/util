package redisCache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
)

type Pagination struct {
	Min       string
	Max       string
	Limit     int64 `json:"limit"`
	Offset    int64 `json:"offset"`
	TotalRows int64 `json:"total_rows"`
}

type SortSet interface {
	// 添加元素
	CacheIndexAdd(ctx context.Context, key string, value []*redis.Z) (result int64, err error)
	// 查询，通过分页
	CacheIndexGetByOrder(ctx context.Context, key string, page *Pagination) (result []string, err error)
	// 查询，通过分数分页
	CacheIndexGetByScore(ctx context.Context, key string, page *Pagination) (result []string, err error)
	// 获取有序集合数量
	CacheGetSortSetCount(ctx context.Context, key string) (result int64, err error)
	// 判断元素是否存在有序集合
	CacheExistsSortSet(ctx context.Context, key, value string) (result bool, err error)
	// 删除有序集合元素
	CacheDelMember(ctx context.Context, key, value string) (err error)
}

type sortSet struct {
	redis redis.Cmdable
}

func NewCacheIndex(redis redis.Cmdable) SortSet {
	return &sortSet{
		redis: redis,
	}
}

func (r *sortSet) CacheIndexAdd(ctx context.Context, key string, value []*redis.Z) (result int64, err error) {
	return r.redis.ZAdd(ctx, key, value...).Result()
}

func (r *sortSet) CacheIndexGetByOrder(ctx context.Context, key string, page *Pagination) (result []string, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return nil, redis.Nil
	}
	var resp []string
	resp, err = r.redis.ZRevRange(ctx, key, page.Offset+1, (page.Offset + page.Limit)).Result()
	page.TotalRows, err = r.CacheGetSortSetCount(ctx, key)
	return resp, err
}

func (r *sortSet) CacheIndexGetByScore(ctx context.Context, key string, page *Pagination) (result []string, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return nil, redis.Nil
	}
	if page.Min == "" {
		page.Min = "0"
	}
	if page.Max == "" {
		page.Max = fmt.Sprintf("%f", math.MaxFloat64)
	}
	var resp []string
	resp, err = r.redis.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
		Min:    page.Min,
		Max:    page.Max,
		Offset: page.Offset,
		Count:  page.Limit,
	}).Result()
	page.TotalRows, err = r.CacheGetSortSetCount(ctx, key)
	return resp, err
}

func (r *sortSet) CacheGetSortSetCount(ctx context.Context, key string) (result int64, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return 0, redis.Nil
	}
	return r.redis.ZCard(ctx, key).Result()
}

func (r *sortSet) CacheExistsSortSet(ctx context.Context, key, value string) (result bool, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return false, redis.Nil
	}
	value = "[" + value
	var count int64
	count, err = r.redis.ZLexCount(ctx, key, value, value).Result()
	if count == 1 {
		return true, nil
	}
	return false, nil
}
func (r *sortSet) CacheDelMember(ctx context.Context, key, value string) (err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return redis.Nil
	}
	_, err = r.redis.ZRem(ctx, key, value).Result()
	return nil
}
