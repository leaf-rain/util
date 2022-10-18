package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func mapIndex(x, y uint64) uint64 {
	return y*1000 + x
}

const (
	MethodMsgIDSize = 16 // 16位字符
)

//go generate ./person.go
func main() {
	var a = 65537
	fmt.Println(uint16(a))
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
