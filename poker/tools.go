package poker

func DefaultValueInt64(length, value int64) []int64 {
	var result = make([]int64, length)
	if value != 0 {
		for i := range result {
			result[i] = value
		}
	}
	return result
}
