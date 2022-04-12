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
	var pusher = kq.NewPusher([]string{"192.168.1.111:9092"}, "stats-topic-t")
	count := rand.Intn(100)
	//准备消息
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
	for i := 0; i < 1000; i++ {
		pusher.Push(string(body))
	}
	cmdline.EnterToContinue()
}

type StatsParam struct {
	Appid        int64   `gorm:"column:appid" db:"appid" json:"appid,omitempty" form:"appid"`                                 //  appid
	Sappid       int64   `gorm:"column:sappid" db:"sappid" json:"sappid,omitempty" form:"sappid"`                             //  子应用id
	Lang         int64   `gorm:"column:lang" db:"lang" json:"lang,omitempty" form:"lang"`                                     //  语言版本
	Pid          int64   `gorm:"column:pid" db:"pid" json:"pid,omitempty" form:"pid"`                                         //  分区id
	Grouping     int64   `gorm:"column:grouping" db:"grouping" json:"grouping,omitempty" form:"grouping"`                     //  分组
	Stypeid      int64   `gorm:"column:stypeid" db:"stypeid" json:"stypeid,omitempty" form:"stypeid"`                         //  统计类型
	Mdate        int64   `gorm:"column:mdate" db:"mdate" json:"mdate,omitempty" form:"mdate"`                                 //  统计时间
	SvalueFloat  float64 `gorm:"column:svalue_float" db:"svalue_float" json:"svalue_float,omitempty" form:"svalue_float"`     //  统计值-float
	SvalueInt    uint64  `gorm:"column:svalue_int" db:"svalue_int" json:"svalue_int,omitempty" form:"svalue_int"`             //  统计值-int
	SvalueString string  `gorm:"column:svalue_string" db:"svalue_string" json:"svalue_string,omitempty" form:"svalue_string"` //  统计值-string
}
