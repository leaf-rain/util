package main

import (
	"fmt"
	"reflect"
	"strings"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

type ttt struct {
	A string
	B string
	C float64
}

var replySpecificKeys = []uint8{1, 2, 3}

func replyGetTy() uint8 {
	if len(replySpecificKeys) == 0 {
		return 0
	}
	replySpecificKeys = append(replySpecificKeys[1:], replySpecificKeys[0])
	return replySpecificKeys[0]
}

func app(list []uint8) {
	replySpecificKeys = append(list, 4, 5, 6)
}

type t1 interface {
	ToString1()
}

// 声明一个空结构体
type cat struct {
	Name string
	// 带有结构体tag的字段
	Type int `json:"type" id:"100"`
}

type t2 interface {
	ToString2()
}

func (c cat) ToString1() {
	fmt.Println(c.Name)
}
func (c cat) ToString2() {
	fmt.Println(c.Name)
}

func main() {
	var req = struct {
		Name string
		Age  int
	}{Name: "张三", Age: 1000}
	var typeOf = reflect.TypeOf(req)
	var valueOf = reflect.ValueOf(req)
	var length = typeOf.NumField()
	var value = make([]interface{}, length*2)
	var name string
	for i := 0; i < length; i++ {
		name = typeOf.Field(i).Name
		value[i*2] = name
		value[i*2+1] = valueOf.FieldByName(name)
	}
	fmt.Printf("%+v", value)
}

func charEncry(name []string, end int) string {
	var length = len(name)
	var encry = "***"
	if length < 5 {
		var right int
		if length >= 3 {
			right = 3
		} else {
			right = length
		}
		return strings.Join(name[:right], "") + encry
	} else {
		return strings.Join(name[:3], "") + encry + strings.Join(name[length-end:], "")
	}
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
