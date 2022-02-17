package redisCache

import (
	"context"
	"fmt"
	rd "github.com/go-redis/redis/v8"
	"github.com/leaf-rain/util/database/redis"
	"testing"
	"time"
)

var client *redis.Client
var ctx context.Context
var sortSetInterface SortSet

func TestMain(m *testing.M) {
	ctx = context.Background()
	var err error
	client, err = redis.NewRedis(redis.Config{
		PoolSize: 5,
		Addr: []string{
			"192.168.1.111:6379",
		},
		DialTimeout: time.Second * 10,
	}, ctx)
	if err != nil {
		panic(err)
	}
	sortSetInterface = NewCacheIndex(client)
	now := time.Now()
	m.Run()
	fmt.Println("函数运行耗时：", time.Since(now))
}

func TestSortSet_CacheSortSet(t *testing.T) {
	var err error
	var result []string
	var resultInt int64
	var key = "test_sort_set"
	if _, err = sortSetInterface.CacheIndexAdd(ctx, key, []*rd.Z{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
	}); err != nil {
		fmt.Errorf("CacheIndexAdd failed, %v", err)
	}
	if result, err = sortSetInterface.CacheIndexGetByOrder(ctx, key, &Pagination{
		Limit:     5,
		Offset:    1,
		TotalRows: 0,
	}); err != nil {
		fmt.Errorf("CacheIndexGetByOrder failed, %v", err)
	} else {
		fmt.Println("CacheIndexGetByOrder success", result)
	}
	if result, err = sortSetInterface.CacheIndexGetByScore(ctx, key, &Pagination{
		Limit:     5,
		Offset:    1,
		TotalRows: 0,
	}); err != nil {
		fmt.Errorf("CacheIndexGetByOrder failed, %v", err)
	} else {
		fmt.Println("CacheIndexGetByOrder success", result)
	}
	if resultInt, err = sortSetInterface.CacheGetSortSetCount(ctx, key); err != nil {
		fmt.Errorf("CacheGetSortSetCount failed, %v", err)
	} else {
		fmt.Println("CacheGetSortSetCount success", resultInt)
	}
}
