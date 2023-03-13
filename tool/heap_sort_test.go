package tool

import "testing"

func Test_heapSort(t *testing.T) {
	var arm Heap = []float64{9.9, 8.8, 7.7, 6.7, 5.5, 1, 2, 3, 4, 0}
	arm.Init()
	t.Logf("%v", arm)
	arm.Push(10.999)
	t.Logf("%v", arm)

}
