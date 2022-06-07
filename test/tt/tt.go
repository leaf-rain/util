package main

import (
	"fmt"
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

func main() {
	for i := int64(1); i <= 15; i++ {
		fmt.Println(i)
	}
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
