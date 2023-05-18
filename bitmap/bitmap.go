package main

import (
	"fmt"
)

type BitMapSwitch struct {
	bitmap []uint64 // 位图数据
}

// 初始化位图开关
func NewBitMapSwitch(numSwitches int) *BitMapSwitch {
	numBits := (numSwitches + 63) / 64 // 计算所需的位图大小
	bitmap := make([]uint64, numBits)
	return &BitMapSwitch{
		bitmap: bitmap,
	}
}

// 打开指定位置的开关
func (b *BitMapSwitch) TurnOn(switchIndex int) {
	if switchIndex >= 0 && switchIndex < len(b.bitmap)*64 {
		wordIndex := switchIndex / 64
		bitIndex := switchIndex % 64
		b.bitmap[wordIndex] |= (1 << bitIndex)
	}
}

// 关闭指定位置的开关
func (b *BitMapSwitch) TurnOff(switchIndex int) {
	if switchIndex >= 0 && switchIndex < len(b.bitmap)*64 {
		wordIndex := switchIndex / 64
		bitIndex := switchIndex % 64
		b.bitmap[wordIndex] &= ^(1 << bitIndex)
	}
}

// 检查指定位置的开关状态
func (b *BitMapSwitch) IsOn(switchIndex int) bool {
	if switchIndex >= 0 && switchIndex < len(b.bitmap)*64 {
		wordIndex := switchIndex / 64
		bitIndex := switchIndex % 64
		return (b.bitmap[wordIndex] & (1 << bitIndex)) != 0
	}
	return false
}

func main() {
	switches := 100 // 开关总数
	bitmapSwitch := NewBitMapSwitch(switches)

	// 打开开关 0、2、4、6
	bitmapSwitch.TurnOn(0)
	bitmapSwitch.TurnOn(2)
	bitmapSwitch.TurnOn(4)
	bitmapSwitch.TurnOn(6)

	// 检查开关状态
	for i := 0; i < switches; i++ {
		fmt.Printf("Switch %d: %v\n", i, bitmapSwitch.IsOn(i))
	}
}
