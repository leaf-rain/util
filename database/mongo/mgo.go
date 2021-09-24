package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
)

type MgoOpt struct {
	Addr            []string `json:"addr" yaml:"addr"` // mongo addr
	Port            string   `json:"port" yaml:"port"`
	MaxConnIdleTime *uint64  `json:"max_conn_idle_time" yaml:"max_conn_idle_time"`
	MaxPoolSize     *uint64  `json:"max_pool_size" yaml:"max_pool_size"`
	MinPoolSize     *uint64  `json:"min_pool_size" yaml:"min_pool_size"`
	ConnectTimeout  *uint64  `json:"connect_timeout" yaml:"connect_timeout"`
	Username        string   `json:"username" yaml:"username"`
	Password        string   `json:"password" yaml:"password"`
}

func NewMgo(opt MgoOpt) (*mongo.Client, context.CancelFunc, error) {
	clientOption := options.Client()
	var addr string

	for i := range opt.Addr {
		addr = addr + "," + opt.Addr[i]
	}
	addr = strings.Trim(addr, ",")
	var url = fmt.Sprintf("mongodb://%s:%s@%s:%s", opt.Username, opt.Password, addr, opt.Port)
	clientOption.ApplyURI(url)
	client, err := mongo.NewClient(clientOption)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = client.Connect(ctx)
	if err != nil {
		return nil, cancel, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, cancel, err
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	return client, cancel, err
}
