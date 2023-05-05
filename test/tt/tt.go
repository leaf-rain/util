package main

import (
	"fmt"
	"sort"
)

type tt struct {
	A int
	B string
}

func main() {
	var a []int64
	v := append([]int64{1, 2, 3}, a...)
	fmt.Println(v)
}

func ta(a *tt) {
	a.A = 1000
}

func Sort(data []int32) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
}

func Pop(data []int32) (int32, []int32) {
	if len(data) == 0 {
		return 0, nil
	}
	if len(data) == 1 {
		return data[0], nil
	}
	result := data[len(data)-1]
	data = data[:len(data)-1]
	return result, data
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
