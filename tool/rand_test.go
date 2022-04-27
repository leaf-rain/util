package tool

import (
	"fmt"
	"testing"
)

func Test_RandInt64(t *testing.T) {
	num := RandInt64(0, 2)
	fmt.Println(num)
}
