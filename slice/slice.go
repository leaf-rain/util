package slice

import "sort"

//通过map键的唯一性去重
func RemoveRepeatedForString(s []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
		if v == "" {
			continue
		}
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// 返回第一列表不包含第二个列表内容
func RemoveStringRepeatedByString(slice, repeatedSlice []string) []string {
	slice = RemoveRepeatedForString(slice)
	result := make([]string, 0)
	var isAdd bool
	for _, item1 := range slice {
		isAdd = true
		for _, item2 := range repeatedSlice {
			if item1 == item2 {
				isAdd = false
			}
		}
		if isAdd {
			result = append(result, item1)
		}
	}
	return result
}

func SortInt32(data []int32) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func SortInt64(data []int64) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func SortUint32(data []uint32) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func Int32ToUint32(data []int32) []uint32 {
	var result = make([]uint32, len(data))
	for index, item := range data {
		result[index] = uint32(item)
	}
	return result
}

func Int32ToInt64(data []int32) []int64 {
	var result = make([]int64, len(data))
	for index, item := range data {
		result[index] = int64(item)
	}
	return result
}

func Int64ToInt32(data []int64) []int32 {
	var result = make([]int32, len(data))
	for index, item := range data {
		result[index] = int32(item)
	}
	return result
}

func Int64ToUint32(data []int64) []uint32 {
	var result = make([]uint32, len(data))
	for index, item := range data {
		result[index] = uint32(item)
	}
	return result
}

func Uint32ToInt32(data []uint32) []int32 {
	var result = make([]int32, len(data))
	for index, item := range data {
		result[index] = int32(item)
	}
	return result
}

func Uint32ToInt64(data []uint32) []int64 {
	var result = make([]int64, len(data))
	for index, item := range data {
		result[index] = int64(item)
	}
	return result
}

func ContainsInt32(data []int32, val int32) bool {
	var result bool
	for _, item := range data {
		if item == val {
			result = true
		}
	}
	return result
}

func ContainsInt64(data []int64, val int64) bool {
	var result bool
	for _, item := range data {
		if item == val {
			result = true
		}
	}
	return result
}

func ContainsUint32(data []uint32, val uint32) bool {
	var result bool
	for _, item := range data {
		if item == val {
			result = true
		}
	}
	return result
}

func LastIndex(index int, length int) int {
	if length <= 0 {
		return -1
	}
	result := index - 1
	if result < 0 {
		return length - 1
	}
	return result
}

func NextIndex(index int, length int) int {
	if length <= 0 {
		return -1
	}
	result := index + 1
	if result >= length {
		return 0
	}
	return result
}
