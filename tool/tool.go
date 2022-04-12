package tool

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
