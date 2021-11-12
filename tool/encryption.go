package tool

import (
	"crypto/sha1"

	"fmt"
)

func GetSha1Str(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	d := s.Sum(nil)
	return fmt.Sprintf("%x", d)
}

func GetLoginPwdStr(str string) string {
	s := sha1.New()
	s.Write([]byte(fmt.Sprintf("%s%s", "13asdf572", str)))
	d := s.Sum(nil)
	return fmt.Sprintf("%x", d)
}

//ClickLoginPwdStr 校验密码
// ss 用户输入密码
// str 数据库用户密码
func ClickLoginPwdStr(ss string, str string) bool {
	s := sha1.New()
	s.Write([]byte(fmt.Sprintf("%s%s", "13asdf572", ss)))
	d := s.Sum(nil)
	return fmt.Sprintf("%x", d) == str
}
