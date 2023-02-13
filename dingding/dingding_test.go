package dingding

import "testing"

func TestDingdingSend(t *testing.T) {
	var svc = NewDingSrv(&Config{
		Secret:    "",
		Urls:      "https://oapi.dingtalk.com/robot/send?access_token=f2a22a6ef7f36d1caa825ec1b95949f8d1d0a6af33037084caf801db3a1321f2",
		AtMobiles: nil,
		AtUserIds: nil,
	})
	result, err := svc.DingdingSend("测试发送")
	if err != nil {
		t.Errorf("failed, err:%v", err)
	} else {
		t.Logf("success, result:%v", result)
	}
}
