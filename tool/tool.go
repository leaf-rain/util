package tool

import (
	"fmt"
	"strconv"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func Uint32ToString(num uint32) string {
	return strconv.FormatUint(uint64(num), 10)
}

//取2位精度并转换成int64
func To2Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
