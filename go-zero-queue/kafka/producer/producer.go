package main

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
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
	var pusher = kq.NewPusher([]string{"192.168.1.111:9092"}, "payment-update-paystatus-topic")
	count := rand.Intn(100)
	// 准备消息
	m := message{
		Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
		Value:   fmt.Sprintf("%d,%d", 100, count),
		Payload: fmt.Sprintf("%d,%d", 100, count),
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
	//cmdline.EnterToContinue()
}
