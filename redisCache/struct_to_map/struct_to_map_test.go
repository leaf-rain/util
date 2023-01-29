package struct_to_map

import (
	"context"
	"fmt"
	"github.com/leaf-rain/util/database/redis"
	"testing"
	"time"
)

var client *redis.Client
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	var err error
	client, err = redis.NewRedis(redis.Config{
		PoolSize: 5,
		Addr: []string{
			"192.168.31.250:6379",
		},
		DialTimeout: time.Second * 10,
	}, ctx)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	m.Run()
	fmt.Println("函数运行耗时：", time.Since(now))
}

var key = "abc"

func TestStorge(t *testing.T) {
	var testStruct = Request{Ping: "ping"}
	err := Storge(ctx, client, key, &testStruct)
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!!")
	}
}

func TestFind(t *testing.T) {
	var result Request
	err := Find(ctx, client, &result, key, []string{})
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!! result :%+v", result)
		t.Logf("success !!! result.ConfInfo :%+v", result.Ping)
	}
}
