package tool

import "testing"

func TestName(t *testing.T) {
	t.Log(GetLoginPwdStr("e10adc3949ba59abbe56e057f20f883e"))
}
