package tool

import (
	"math"
	"sync"
)

// 本例为最大堆
// 最小堆只需要修改less函数即可
type Heap struct {
	data []*HeapInfo
	lock *sync.RWMutex
}

type HeapInfo struct {
	Id    int64
	Score int64
}

func NewHeap() *Heap {
	return &Heap{lock: new(sync.RWMutex)}
}

func (h Heap) init() {
	n := len(h.data)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x *HeapInfo) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, item := range h.data {
		if item.Id == x.Id {
			item.Score = x.Score
			h.init()
			return
		}
	}
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) RemoveByIndex(i int) (*HeapInfo, bool) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if i < 0 || i > len(h.data)-1 {
		return nil, false
	}
	n := len(h.data) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (h.data)[n]
	h.data = (h.data)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (h.data)[i].Score > (h.data)[(i-1)/2].Score {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) RemoveById(id int64) bool {
	h.SetScore(id, math.MinInt64)
	h.lock.Lock()
	defer h.lock.Unlock()
	for index, _ := range h.data {
		if h.data[index].Score == math.MinInt64 {
			h.data = h.data[:index]
			return true
		}
	}
	return false
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() *HeapInfo {
	h.lock.Lock()
	defer h.lock.Unlock()
	n := len(h.data) - 1
	if n < 0 {
		return nil
	}
	h.swap(0, n)
	x := (h.data)[n]
	h.data = (h.data)[0:n]
	h.down(0)
	return x
}

// 修改某个id的分数
func (h *Heap) SetScore(id, score int64) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, item := range h.data {
		if item.Id == id {
			item.Score = score
			h.init()
			break
		}
	}
}

func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h.data) {
			break // i已经是叶子结点了
		}
		j := l
		if r := l + 1; r < len(h.data) && h.less(r, l) {
			j = r // 右孩子
		}
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}

func (h Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h Heap) less(i, j int) bool {
	return h.data[i].Score >= h.data[j].Score
}
