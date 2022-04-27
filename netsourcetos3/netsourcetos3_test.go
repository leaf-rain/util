package netsourcetos3

import (
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

var conf = S3Config{
	AccessKey: "AKIA5DQBT52JAQPARCUV",
	SecretKey: "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",
	Token:     "",
	EndPoint:  "s3.amazonaws.com",
	Region:    "ap-south-1",
}

var url = "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"

func TestMain(m *testing.M) {
	InitS3(conf)
	now := time.Now()
	m.Run()
	logx.Infof("程序执行耗时：%v", time.Since(now))
}

func TestNetsourcetos3ByGet(t *testing.T) {
	result, err := Netsourcetos3ByGet(url, "touchatimg143435-prod", "public/album/", func(str string) string {
		return "https://img.tapmechat.com/album/" + str
	})
	if err != nil {
		t.Errorf(" failed, err:%v", err)
		return
	}
	t.Logf("success, result:%+v", result)
}
