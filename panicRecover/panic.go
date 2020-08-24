package panicRecover

import (
	"fmt"
	"github.com/leaf-rain/log"
	"runtime"
	"go.uber.org/zap"
)

// RecoverPanic 恢复panic
func RecoverPanic() {
	err := recover()
	if err != nil {
		log.Error("panic,err:%v, stack:%v", zap.Any("panic", err), zap.String("stack", GetStackInfo()))
	}
}

// GetStackInfo 获取Panic堆栈信息
func GetStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return fmt.Sprintf("%s", buf[:n])
}

