package tool

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrOutOfMaps = errors.New("位图越界")
)

type Maps struct {
	width uint64
	high  uint64
	lock  *sync.RWMutex
	*BitMap
}

func NewMaps(width, high uint64) *Maps {
	return &Maps{
		BitMap: NewBitMap(width * high),
		width:  width,
		high:   high,
		lock:   new(sync.RWMutex),
	}
}

func (m *Maps) Add(x, y, width, high uint64) error {
	if x+width > m.width || y+high > m.high {
		return ErrOutOfMaps
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	for i1 := x; i1 < x+width; i1++ {
		for i2 := y; i2 < y+high; i2++ {
			var v = m.value(i1, i2)
			m.BitMap.Add(v)
			fmt.Println(v)
		}
	}
	return nil
}

func (m *Maps) Remove(x, y, width, high uint64) error {
	if x+width > m.width || y+high > m.high {
		return ErrOutOfMaps
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	for i1 := x; i1 <= x+width; i1++ {
		for i2 := y; i2 <= y+high; i2++ {
			m.BitMap.Remove(m.value(i1, i2))
		}
	}
	return nil
}

func (m *Maps) value(x, y uint64) uint64 {
	return y*m.high + x
}

//位图
type BitMap struct {
	bits []byte
	max  uint64
}

//初始化一个BitMap
//一个byte有8位,可代表8个数字,取余后加1为存放最大数所需的容量
func NewBitMap(max uint64) *BitMap {
	bits := make([]byte, (max>>3)+1)
	return &BitMap{bits: bits, max: max}
}

//添加一个数字到位图
//计算添加数字在数组中的索引index,一个索引可以存放8个数字
//计算存放到索引下的第几个位置,一共0-7个位置
//原索引下的内容与1左移到指定位置后做或运算
func (b *BitMap) Add(num uint64) {
	index := num >> 3
	pos := num & 0x07
	b.bits[index] |= 1 << pos
}

//判断一个数字是否在位图
//找到数字所在的位置,然后做与运算
func (b *BitMap) IsExist(num uint64) bool {
	index := num >> 3
	pos := num & 0x07
	return b.bits[index]&(1<<pos) != 0
}

//删除一个数字在位图
//找到数字所在的位置取反,然后与索引下的数字做与运算
func (b *BitMap) Remove(num uint64) {
	index := num >> 3
	pos := num & 0x07
	b.bits[index] = b.bits[index] & ^(1 << pos)
}

//位图的最大数字
func (b *BitMap) Max() uint64 {
	return b.max
}

func (b *BitMap) String() string {
	return fmt.Sprint(b.bits)
}
