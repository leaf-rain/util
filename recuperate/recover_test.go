package recuperate

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGoroutineNotPanic(t *testing.T) {
	var handlers []func() error

	handlers = append(handlers, func() error {
		time.Sleep(time.Second * 2)
		panic("手动错误")
	})

	err := GoroutineNotPanic(handlers...)
	for _, item := range err {
		t.Log("err: \n", item)
	}
	t.Log("不影响程序运行！！！")
}

func TestPanicGroup(t *testing.T) {
	NewPanicGroup().Go(func() {
		defer func() {
			fmt.Println("-->1")
		}()
		time.Sleep(time.Second * 2)
		panic("手动错误1")
	}).Go(func() {
		time.Sleep(time.Second * 3)
		defer func() {
			fmt.Println("-->2")
		}()
		time.Sleep(time.Second * 2)
		panic("手动错误2")
	}).Go(func() {
		time.Sleep(time.Second * 3)
		defer func() {
			fmt.Println("-->3")
		}()
		time.Sleep(time.Second * 2)
		panic("手动错误3")
	}).Wait(context.Background())
	t.Log("不影响程序运行！！！")
}
