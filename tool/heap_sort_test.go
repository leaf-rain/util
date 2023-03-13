package tool

import "testing"

func Test_heapSort(t *testing.T) {
	var arm Heap = []HeapInfo{
		{
			Score: 9,
		},
		{
			Score: 8,
		},
		{
			Score: 1000,
		},
		{
			Score: 1,
		},
	}
	arm.Init()
	t.Logf("%v", arm)
	arm.Push(HeapInfo{
		Score: 1000000000,
		Param: nil,
	})
	t.Logf("%v", arm)

}
