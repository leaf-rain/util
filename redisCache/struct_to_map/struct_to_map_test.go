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

type ts struct {
	Name string `json:"Name,omitempty"`
	Age  int64  `json:"Age,omitempty"`
}

type tts struct {
	Name   string              `json:"Name,omitempty"`
	Age    int64               `json:"Age,omitempty"`
	Values *ts                 `json:"Values,omitempty"`
	Slice  []int               `json:"Slice"`
	Map    map[int]interface{} `json:"Map"`
}

var key = "abc"

func TestStorge(t *testing.T) {
	var testStruct = tts{
		Name:  "张三",
		Age:   30,
		Slice: []int{1, 2, 3},
	}
	testStruct.Values = &ts{
		Name: "李四",
		Age:  40,
	}
	testStruct.Map = make(map[int]interface{})
	testStruct.Map[1] = 2
	testStruct.Map[2] = 3
	err := Storge(ctx, client, key, testStruct)
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!!")
	}
}

func TestFind(t *testing.T) {
	var result tts
	err := Find(ctx, client, &result, key, []string{"Slice"})
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!! result :%+v", result)
		t.Logf("success !!! result.Slice :%+v", result.Slice)
		t.Logf("success !!! result.Values :%+v", result.Values)
	}
}
