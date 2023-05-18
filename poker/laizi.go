package poker

import (
	"sort"
)

func (p *Poker) GetLaizi() []int64 {
	return p.laizi
}

func (p *Poker) SetLaizi(data []int64) {
	p.laizi = data
	for i1 := range data {
		for i2 := range p.baseCards {
			if p.baseCards[i2].Value == data[i1] {
				p.baseCards[i2].IsLaizi = true
			}
		}
	}
	return
}

func (p *Poker) AppendLaizi(data []int64) {
	p.laizi = append(p.laizi, data...)
	sort.SliceIsSorted(p.laizi, func(i, j int) bool { // 排序
		return p.laizi[i] < p.laizi[j]
	})
	for i1 := range data {
		for i2 := range p.baseCards {
			if p.baseCards[i2].Value == data[i1] {
				p.baseCards[i2].IsLaizi = true
			}
		}
	}
	return
}

func (p *Poker) IsLaizi(card int64) bool {
	for i := range p.laizi {
		if p.laizi[i] == card {
			return true
		}
	}
	return false
}

// 自动生成癞子牌
func (p *Poker) AutoLaizi(num int64) { // num:癞子数量
	cards := p.RandomGetDifferent(p.baseCards, int(num))
	p.laiNum = int64(len(cards))
	p.laizi = make([]int64, p.laiNum)
	for i := range cards {
		p.laizi[i] = cards[i].Value
		for i2 := range p.baseCards {
			if p.baseCards[i2].Value == cards[i].Value {
				p.baseCards[i2].IsLaizi = true
			}
		}
	}
}

// 包含癞子数量
func (p *Poker) ContainLaizi(cards interface{}) int64 {
	if p.laiNum == 0 {
		return 0
	}
	var (
		numCards int64
	)
	switch cards.(type) {
	case Card:
		card := cards.(Card)
		if p.IsLaizi(card.Value) {
			numCards += 1
		}
	case *Card:
		card := cards.(*Card)
		if p.IsLaizi(card.Value) {
			numCards += 1
		}
	case []*Card:
		cs := cards.([]*Card)
		for _, c := range cs {
			if p.IsLaizi(c.Value) {
				numCards += 1
			}
		}
	default:
	}
	return numCards
}
