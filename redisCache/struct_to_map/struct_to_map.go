package struct_to_map

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

var (
	ErrParam    = errors.New("参数不正确")
	ErrNotFound = errors.New("查询结果为0")
)

func Storge(ctx context.Context, redisClient redis.Cmdable, key string, data ...interface{}) error {
	if len(data) == 1 {
		data = Encode(data[0])
	}
	if len(data) == 0 {
		return ErrParam
	}
	return redisClient.HMSet(ctx, key, data...).Err()
}

func Find(ctx context.Context, redisClient redis.Cmdable, dst interface{}, key string, fields []string) error {
	var rdResult map[string]string
	var err error
	if len(fields) == 0 { // 查询字段为0时默认查询全字段
		rdResult, err = redisClient.HGetAll(ctx, key).Result()
		if err != nil {
			return err
		}
	} else {
		var rdResultSlice []interface{}
		rdResultSlice, err = redisClient.HMGet(ctx, key, fields...).Result()
		if err != nil {
			return err
		}
		rdResult = make(map[string]string, len(rdResultSlice))
		var ok bool
		for index, item := range rdResultSlice {
			if item == nil {
				continue
			}
			if _, ok = item.(string); ok {
				rdResult[fields[index]] = item.(string)
			}
		}
	}
	if len(rdResult) == 0 {
		return ErrNotFound
	}
	return Decode(dst, rdResult)
}
