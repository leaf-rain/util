package main

import (
	"fmt"
	"github.com/shima-park/agollo"
	"os"
	"strings"
)

func main() {
	app, _ := agollo.New("192.168.2.45:8080", "vchat", agollo.AutoFetchOnCacheMiss())
	var nps = app.Get("finance")
	fmt.Println("get nps success, nps:", nps)
	npList := strings.Split(nps, ";")
	fmt.Println("nplist:", npList)

	// 通过默认根目录下的app.properties初始化agollo
	agollo.Init(
		"192.168.2.45:8080",
		"vchat",
		agollo.AutoFetchOnCacheMiss(),
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))), // 打印日志信息
		agollo.PreloadNamespaces(npList...),                                 // 预先加载的namespace列表，如果是通过配置启动，会在app.properties配置的基础上追加
		agollo.AutoFetchOnCacheMiss(),                                       // 在配置未找到时，去apollo的带缓存的获取配置接口，获取配置
		agollo.FailTolerantOnBackupExists(),                                 // 在连接apollo失败时，如果在配置的目录下存在.agollo备份配置，会读取备份在服务器无法连接的情况下
	)

	// 获取namespace下的所有配置项
	fmt.Println("Configuration of the namespace:", agollo.GetNameSpace("mihua.common.yml"))
	//fmt.Println(agollo.Get("test"))
	// 如果想监听并同步服务器配置变化，启动apollo长轮训
	// 返回一个期间发生错误的error channel,按照需要去处理
	errorCh := agollo.Start()
	defer agollo.Stop()

	watchCh := agollo.Watch()

	for {
		select {
		case err := <-errorCh:
			fmt.Println("Error:", err)
		case resp := <-watchCh:
			fmt.Println("Watch Apollo:", resp)
		}
	}
}
