package tool

import (
	"regexp"
	"strconv"
)

func IsThPhone(phone string) bool {
	re := regexp.MustCompile(`^(?:66)?0?(\d{9})$`)
	return re.MatchString(phone)
}

func FixThPhone(phone string) string {
	s0x := regexp.MustCompile(`^0?(\d{9})$`)
	fixreg := regexp.MustCompile(`^0?`)
	if s0x.MatchString(phone) {
		//0xxxx的手机号
		fix := fixreg.ReplaceAllString(phone, "66")
		return fix
	} else {
		//66xxx的手机号 无需修改
		return phone
	}
}

func GetVIPLevel(VIP string) (bool, int) {
	lv := regexp.MustCompile(`^VIP达到(\d+)级$`)
	if lv.MatchString(VIP) {
		ll := lv.FindStringSubmatch(VIP)
		//fmt.Println(ll)
		if len(ll) == 2 {
			if l, err := strconv.Atoi(ll[1]); err != nil {
				return false, 0
			} else {
				return true, l
			}
		}
		return false, 0
	} else {
		return false, 0
	}
}

func GetVIPUpdate(VIP string) (bool, int) {
	lv := regexp.MustCompile(`^VIP升级.*=(\d+)\)$`)
	if lv.MatchString(VIP) {
		ll := lv.FindStringSubmatch(VIP)
		if len(ll) == 2 {
			if l, err := strconv.Atoi(ll[1]); err != nil {
				return false, 0
			} else {
				return true, l
			}
		}
		return false, 0
	} else {
		return false, 0
	}
}
