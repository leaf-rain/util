package recuperate

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const (
	defaultStackSize = 4096
)

func Recover() {
	if err := recover(); err != nil {
		log.Printf("panic recover, err: %v", err)
		debug.PrintStack()
	}
}

func GoroutineNotPanic(handlers ...func() error) (err []error) {
	var wg sync.WaitGroup
	for _, f := range handlers {
		wg.Add(1)
		go func(handler func() error) {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("panic recover, err: %v", err)
					buf := make([]byte, 64<<10) //64*2^10, 64KB
					buf = buf[:runtime.Stack(buf, false)]
					err = append(err, fmt.Errorf("panic recovered: %s\n %s", e, buf))
				}
				wg.Done()
			}()
			e := handler()
			err = append(err, e)
		}(f)
	}
	wg.Wait()
	return
}

func GoSafe(f func()) {
	go func() {
		defer Recover()
		f()
	}()
}

// getCurrentGoroutineStack 获取当前Goroutine的调用栈，便于排查panic异常
func getCurrentGoroutineStack() string {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

type Panic struct {
	R     interface{} // recover() 返回值
	Stack []byte      // 当时的调用栈
}

func (p Panic) String() string {
	return fmt.Sprintf("%v\n%s", p.R, p.Stack)
}

type PanicGroup struct {
	panics chan Panic // 协程 panic 通知信道
	dones  chan int   // 协程完成通知信道
	jobN   int32      // 协程并发数量
}

func NewPanicGroup() *PanicGroup {
	return &PanicGroup{
		panics: make(chan Panic, 8),
		dones:  make(chan int, 8),
	}
}

func (g *PanicGroup) Go(f func()) *PanicGroup {
	atomic.AddInt32(&g.jobN, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				g.panics <- Panic{R: r, Stack: debug.Stack()}
			}
			g.dones <- 1
		}()
		f()
	}()

	return g // 方便链式调用
}

func (g *PanicGroup) Wait(ctx context.Context) error {
	for {
		select {
		case <-g.dones:
			if atomic.AddInt32(&g.jobN, -1) == 0 {
				return nil
			}
		case p := <-g.panics:
			fmt.Println(p.String())
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
