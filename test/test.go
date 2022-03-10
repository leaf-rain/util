package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.111:2379"},
		DialTimeout: time.Second * 10,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "test_key", "test_value")
	defer cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "q1mi")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get from etcd failed, err:%v\n", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}
}
