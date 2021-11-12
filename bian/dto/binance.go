package dto

import (
	"github.com/leaf-rain/util/bian/consts"
	"github.com/leaf-rain/util/tool"
	"net/url"
)

type Parameters struct {
	RecvWindow int64 `json:"recvWindow"` // 时间差
	Timestamp  int64 `json:"timestamp"`  // 时间戳
}

func (a *Parameters) ToString() string {
	if a.RecvWindow <= 0 {
		a.RecvWindow = consts.RecvWindow
	}
	if a.Timestamp <= 0 {
		a.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	values.Add("recvWindow", tool.Int64ToString(a.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(a.Timestamp))
	return values.Encode()
}
