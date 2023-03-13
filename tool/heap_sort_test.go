package tool

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	var arm *Heap = NewHeap()
	arm.Push(&HeapInfo{
		Id:    1000,
		Score: 9,
	})
	arm.Push(&HeapInfo{
		Id:    1000,
		Score: 8,
	})
	arm.RemoveById(1000)
	fmt.Println(arm.Pop())
}
