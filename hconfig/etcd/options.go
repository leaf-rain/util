package etcd

import (
	"golang.org/x/net/context"
)

type Option func(*options)

type options struct {
	ctx   context.Context
	root  string
	paths []string
}

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithRoot(prefix string) Option {
	return func(o *options) {
		o.root = prefix
	}
}

func WithPaths(path ...string) Option {
	return func(o *options) {
		o.paths = path
	}
}

func NewOptions(opts ...Option) *options {
	options := &options{
		ctx:   context.Background(),
		root:  "/config",
		paths: []string{"test"},
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
