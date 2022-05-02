package main

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// message structure
type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	var pusher = kq.NewPusher([]string{"192.168.1.1111:9092"}, "sysconfig-topic")
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for round := 0; round < 10; round++ {
		select {
		case <-ticker.C:
			count := rand.Intn(100)
			// 准备消息
			m := message{
				Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
				Value:   fmt.Sprintf("%d,%d", round, count),
				Payload: fmt.Sprintf("%d,%d", round, count),
			}
			body, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body))
			// push to kafka broker
			if err := pusher.Push(string(body)); err != nil {
				log.Fatal(err)
			}
		}
	}
	cmdline.EnterToContinue()
}
