package redisCache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	rd "github.com/leaf-rain/util/database/redis"
	"time"
)

/* 该延时队列为简单实现，暂不支持重试和死信 */

type DelayedQueue interface {
	Put(ctx context.Context, key string, z []*redis.Z, ttl time.Duration) error
	DealWithOnce(ctx context.Context, key string, min, max float64, ttl time.Duration) error
	DealWithMultiple(ctx context.Context, key string, min, max float64, ttl time.Duration) error
}

type DealWith interface {
	DealWithOnce(z redis.Z) error
	DealWithMultiple(z []redis.Z) error
}

type defaultQueue struct {
	rd       *rd.Client
	lock     RedisLock
	dealWith DealWith
}

func NewDelayedQueue(rd *rd.Client, fun DealWith) DelayedQueue {
	return &defaultQueue{
		rd:       rd,
		lock:     NewRedisLock(rd),
		dealWith: fun,
	}
}

func (d defaultQueue) Put(ctx context.Context, key string, z []*redis.Z, ttl time.Duration) error {
	err := d.rd.ZAdd(ctx, key, z...).Err()
	if err != nil {
		return err
	}
	_ = d.rd.Expire(ctx, key, ttl) // 不一定非要设定成功
	return err
}

func (d defaultQueue) DealWithOnce(ctx context.Context, key string, min, max float64, ttl time.Duration) error {
	lock, _, err := d.lock.TryLock(ctx, key, 0, ttl)
	if err != nil {
		return err
	}
	if !lock {
		//logx.WithContext(ctx).Infof("[DelayedQueue] try lock failed. key:%s", key)
		return nil
	}
	var z []redis.Z
	z, err = d.rd.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    fmt.Sprintf("%f", min),
		Max:    fmt.Sprintf("%f", max),
		Offset: 0,
		Count:  1,
	}).Result()
	if len(z) == 0 {
		//logx.WithContext(ctx).Infof("[DelayedQueue] ZRangeByScoreWithScores z length zero. key:%s", key)
		return nil
	}
	err = d.dealWith.DealWithOnce(z[0])
	if err == nil {
		return d.rd.ZRem(ctx, key, z[0].Member).Err()
	} else {
		return err
	}
}

func (d defaultQueue) DealWithMultiple(ctx context.Context, key string, min, max float64, ttl time.Duration) error {
	lock, _, err := d.lock.TryLock(ctx, key, 0, ttl)
	if err != nil {
		return err
	}
	if !lock {
		//logx.WithContext(ctx).Infof("[DelayedQueue] try lock failed. key:%s", key)
		return nil
	}
	var z []redis.Z
	z, err = d.rd.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    fmt.Sprintf("%f", min),
		Max:    fmt.Sprintf("%f", max),
		Offset: 0,
		Count:  0,
	}).Result()
	if len(z) == 0 {
		//logx.WithContext(ctx).Infof("[DelayedQueue] ZRangeByScoreWithScores z length zero. key:%s", key)
		return nil
	}
	err = d.dealWith.DealWithMultiple(z)
	if err == nil {
		var member = make([]interface{}, len(z))
		for i := range z {
			member[i] = z[i].Member
		}
		return d.rd.ZRem(ctx, key, member...).Err()
	} else {
		return err
	}
}
