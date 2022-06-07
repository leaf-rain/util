package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

func main() {
	tw, err := collection.NewTimingWheel(time.Second, 300, func(key, value interface{}) {
		fmt.Println(key)
		fmt.Println(value)
	})
	if err != nil {
		logrus.Fatalf("init timingwheel error:%v", err)
		return
	}
	defer tw.Stop()
	//var wg = sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	defer wg.Add(-1)
	//
	//}()
	//wg.Wait()
	for i := 0; i < 10; i++ {
		tw.SetTimer(i, struct {
			A int64
			B int64
		}{A: 1, B: 2}, time.Second*time.Duration(i))
		//time.Sleep(time.Second)
	}
	fmt.Println("========")
	time.Sleep(time.Second * 100)
}
