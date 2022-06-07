package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func main() {
	logx.MustSetup(logx.LogConf{
		ServiceName: "测试",
		Mode:        "file",
		Path:        "./logx/logs",
		KeepDays:    3,
	})
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		logx.Infof("time:%d", time.Now().Unix())
	}
}
