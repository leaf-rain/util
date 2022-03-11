package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"strconv"
	"sync"
)

func main() {
	var c kq.KqConf = kq.KqConf{
		ServiceConf: service.ServiceConf{
			Name: "Test",
			Log: logx.LogConf{
				ServiceName: "Test_kq",
				Mode:        "console",
				Path:        "",
			},
			Mode: "",
		},
		Brokers:    []string{"192.168.2.62:9092"},
		Group:      "payment-update-paystatus-group",
		Topic:      "payment-update-paystatus-topic",
		Offset:     "first",
		Consumers:  1,
		Processors: 1,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var con = consumer{ctx: context.Background(), name: "consume_" + strconv.FormatInt(int64(i), 10)}
			q := kq.MustNewQueue(c, &con)
			defer q.Stop()
			q.Start()
		}(i)
	}
	wg.Wait()
}

type consumer struct {
	ctx  context.Context
	name string
}

func (c consumer) Consume(k, v string) error {
	fmt.Printf("%s,=======> %s => %s\n", c.name, k, v)
	return nil
}
