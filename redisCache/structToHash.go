package redisCache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"reflect"
)

var ErrUnknownType = errors.New("未知数据类型")

func StructToRedisHash(ctx context.Context, cli redis.Cmdable, data interface{}) error {
	var typeOf = reflect.TypeOf(data)
	var valueOf = reflect.ValueOf(data)
	var length = typeOf.NumField()
	var filter = make([]interface{}, length*2)
	for i := 0; i < length; i++ {
		filter[i*2] = typeOf.Field(i).Name
		var value = valueOf.FieldByName(typeOf.Field(i).Name)
		switch value.Interface().(type) { // redis只支持存储基本类型数据，复杂类型数据需要序列化成[]byte类型
		case bool,
			int,
			int8,
			int16,
			int32,
			int64,
			uint,
			uint8,
			uint16,
			uint32,
			uint64,
			float32,
			float64,
			complex64,
			complex128,
			[]byte,
			string:
			break
		default:
			panic(ErrUnknownType)
		}
		filter[i*2+1] = value.Interface()
	}
	return cli.HMSet(ctx, "test", filter...).Err()
}
