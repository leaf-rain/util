package tool

import (
	"fmt"
	"testing"
)

func TestNewMaps(t *testing.T) {
	var x, y uint64 = 3, 0
	var width, high uint64 = 4, 3

	var m = NewMaps(10, 10)
	m.Add(x, y, width, high)
	fmt.Println(m.IsExist(m.value(3, 0)))
	fmt.Println(m.IsExist(m.value(3, 3)))
	fmt.Println(m.IsExist(m.value(6, 0)))
	fmt.Println(m.IsExist(m.value(6, 2)))
	fmt.Println(m.IsExist(m.value(0, 0)))
	fmt.Println(m.String())
}
