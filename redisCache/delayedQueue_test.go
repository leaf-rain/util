package redisCache

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

type dealwith struct {
}

func (dealwith) DealWithOnce(z redis.Z) error {
	logx.Infof("%+v", z)
	logx.Infof("%v", time.Now().Unix())
	return nil
}
func (dealwith) DealWithMultiple(z []redis.Z) error {
	for i := range z {
		logx.Infof("%+v", z[i])
		logx.Infof("%v", time.Now().Unix())
	}
	return nil
}

func TestNewDelayedQueue(t *testing.T) {
	queue := NewDelayedQueue(client, dealwith{})
	var key = "text_delaye_queue"
	go func() {
		var now = time.Now()
		for i := 10; i < 100; i++ {
			queue.Put(ctx, key, []*redis.Z{{
				Score:  float64(now.Unix()) + float64(i),
				Member: now.Unix() + int64(i),
			}}, time.Hour*6)
		}
	}()
	go func() {
		for {
			var now = time.Now()
			queue.DealWithMultiple(ctx, key, float64(now.Unix()), time.Second*3)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Hour)
}
