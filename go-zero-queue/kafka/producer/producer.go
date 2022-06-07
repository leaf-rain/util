package main

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"log"
)

// message structure
type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func main() {
	var pusher = kq.NewPusher([]string{"192.168.1.111:9092"}, "sysconfig-topic")
	// push to kafka broker
	if err := pusher.Push("Router"); err != nil {
		log.Fatal(err)
	}
	cmdline.EnterToContinue()
}
