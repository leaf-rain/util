package redisCache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Pagination struct {
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
	TotalRows int `json:"total_rows"`
}

type SortSet interface {
	CacheIndexAdd(ctx context.Context, key string, value []*redis.Z) (result int64, err error)
	CacheIndexGetByOrder(ctx context.Context, key string, page *Pagination) (result []int64, err error)
	CacheIndexGetByScore(ctx context.Context, key string, page *Pagination) (result []int64, err error)
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

func (r *sortSet) CacheIndexGetByOrder(ctx context.Context, key string, page *Pagination) (result []int64, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return nil, redis.Nil
	}
	var resp []string
	resp, err = r.redis.ZRevRange(ctx, key, int64(page.Offset), int64(page.Offset+page.Limit-1)).Result()
	for i := range resp {
		num, err := strconv.ParseInt(resp[i], 10, 64)
		if err == nil {
			result = append(result, num)
		}
	}
	return result, err
}

func (r *sortSet) CacheIndexGetByScore(ctx context.Context, key string, page *Pagination) (result []int64, err error) {
	if r.redis.Exists(ctx, key).Val() != 1 {
		return nil, redis.Nil
	}
	var resp []string
	resp, err = r.redis.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Offset: int64(page.Offset),
		Count:  int64(page.Limit),
	}).Result()
	for i := range resp {
		num, err := strconv.ParseInt(resp[i], 10, 64)
		if err == nil {
			result = append(result, num)
		}
	}
	return result, err
}
