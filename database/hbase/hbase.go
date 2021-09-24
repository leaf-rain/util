package hbase

import "github.com/tsuna/gohbase"

type HBaseConf struct {
	addr string
}

func NewHbase(str *HBaseConf) gohbase.Client {
	return gohbase.NewClient(str.addr)
}
