package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func mapIndex(x, y uint64) uint64 {
	return y*1000 + x
}

const (
	MethodMsgIDSize = 16 // 16位字符
)

//go generate ./person.go
func main() {
	fmt.Println(100 / 256)
}

func Camel2Case2(name string) string {
	var result = ""
	var last = -1
	var length = len(name)
	for i, r := range name {
		value := string(r)
		if unicode.IsUpper(r) {
			if i > 0 && unicode.IsUpper(rune(name[i-1])) && i+1 < length && !unicode.IsUpper(rune(name[i+1])) {
				result += "_" + strings.ToLower(value)
			} else {
				if last+1 == i {
					result += strings.ToLower(value)
				} else {
					result += "_" + strings.ToLower(value)
				}
			}
			last = i
		} else {
			result += value
		}
	}
	return result
}

func Camel2Case(name string) string {
	var index = -1
	var slice = make([]string, 0)
	var isUpper = make([]int, 0)
	for i, r := range name {
		value := string(r)
		if unicode.IsUpper(r) {
			if i > 0 && unicode.IsUpper(rune(name[i-1])) {
				if index < 0 {
					index = 0
				}
				slice[index] = slice[index] + value
			} else {
				index++
				isUpper = append(isUpper, index)
				slice = append(slice, value)
			}
		} else {
			index++
			slice = append(slice, value)
		}
	}
	fmt.Println(slice)
	fmt.Println(isUpper)
	return strings.Join(slice, "")
}

func search(nums []int, target int) int {
	var l = 0
	var r = len(nums) - 1
	var i int
	for l <= r {
		i = l + (r-l)/2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			l = i + 1
		} else if nums[i] > target {
			r = i - 1
		}
	}
	return -1
}

func GenerateRandInt(min, max uint32) uint32 {
	if max < min {
		return 0
	}
	rand.Seed(time.Now().UnixNano()) //随机种子
	result := rand.Intn(int(max-min)) + int(min)
	return uint32(result)

}

// 蛇形转驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func StringAppendUnRepeat(data []string, item string) {
	var needAppend = true
	for i := range data {
		if data[i] == item {
			needAppend = false
		}
	}
	if needAppend {
		data = append(data, item)
	}
}

func ByteToArgumentBody(req []byte) *ArgumentBody {
	if len(req) < MethodMsgIDSize {
		return nil
	}
	var requestBody = &ArgumentBody{}
	requestBody.Id = string(req[:MethodMsgIDSize])
	requestBody.Body = req[MethodMsgIDSize:]
	return requestBody
}

func ArgumentBodyToByte(req *ArgumentBody) []byte {
	if len(req.Id) == 0 {
		return nil
	}
	var body []byte
	if req.Any != nil {
		if msg, ok := req.Any.(proto.Message); !ok {
			return nil
		} else {
			body, _ = proto.Marshal(msg)
		}
	} else {
		body = req.Body
	}
	var result = make([]byte, MethodMsgIDSize+len(body))
	for index, item := range req.Id {
		result[index] = byte(item)
	}
	copy(result[MethodMsgIDSize:], body)
	return result
}

func CheckRepeatForList(l1, l2 []uint64) ([]uint64, float64) {
	var resultList = make([]uint64, 0)
	var resultNum float64
	var length1 = len(l1)
	var length2 = len(l2)
	if length1 <= 1024 || length2 <= 1024 { // golang中小于1024的话，直接下标遍历数组，会优于map
		for i1 := range l1 {
			for i2 := range l2 {
				if l1[i1] == l2[i2] {
					resultList = append(resultList, l1[i1])
					resultNum += 1
				}
			}
		}
		return resultList, resultNum
	}
	var m = make(map[uint64]struct{})
	if length1 > length2 {
		for i := range l1 {
			m[l1[i]] = struct{}{}
		}
		for i := range l2 {
			if _, ok := m[l2[i]]; ok {
				resultList = append(resultList, l2[i])
				resultNum += 1
			}
		}
		return resultList, resultNum
	} else {
		for i := range l2 {
			m[l2[i]] = struct{}{}
		}
		for i := range l1 {
			if _, ok := m[l1[i]]; ok {
				resultList = append(resultList, l1[i])
				resultNum += 1
			}
		}
		return resultList, resultNum
	}
}
