package tool

import (
	"github.com/zeromicro/go-zero/core/logx"
	"runtime"
	"runtime/debug"
)

const (
	defaultStackSize = 4096
)

func Recover() {
	if err := recover(); err != nil {
		logx.Errorf("panic recover, err: %v", err)
		debug.PrintStack()
	}
}

// getCurrentGoroutineStack 获取当前Goroutine的调用栈，便于排查panic异常
func getCurrentGoroutineStack() string {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
