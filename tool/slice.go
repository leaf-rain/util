package tool

func IntInSlice(num int, slice []int) (exists bool) {
	if len(slice) == 0 {
		return false
	}
	for i := range slice {
		if slice[i] == num {
			exists = true
			return exists
		}
	}
	return exists
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 元素去重
func RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return RemoveRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return RemoveRepByMap(slc)
	}
}

//通过map键的唯一性去重
func RemoveRepeatedForString(s []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool) //map的值不重要
	for _, v := range s {
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
