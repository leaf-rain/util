package redis

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type un int64

type st struct {
	Name string `redis:"Name"`
	Age  int64  `redis:"Age"`
	Un   un     `redis:"Un"`
}

var ErrUnknownType = errors.New("未知类型")

func TestNew(t *testing.T) {
	var ctx = context.Background()
	cli, err := NewRedis(Config{
		PoolSize: 5,
		Addr: []string{
			"127.0.0.1:6379",
		},
		DialTimeout: time.Second * 10,
	}, ctx)
	if err != nil {
		panic(err)
	}
	var data = st{
		Name: "张三",
		Age:  10,
		Un:   20,
	}
	var typeOf = reflect.TypeOf(data)
	var valueOf = reflect.ValueOf(data)
	var length = typeOf.NumField()
	var filter = make([]interface{}, length*2)
	for i := 0; i < length; i++ {
		filter[i*2] = typeOf.Field(i).Name
		value := valueOf.FieldByName(typeOf.Field(i).Name)
		switch value.Interface().(type) {
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
	err = cli.HMSet(ctx, "test", filter...).Err()
	if err != nil {
		panic(err)
	}
	var data2 st
	err = cli.HGetAll(ctx, "test").Scan(&data2)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)
}
