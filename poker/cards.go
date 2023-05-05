package poker

import (
	"math/rand"
	"sort"
	"time"
)

// Card 扑克牌对象，牌值 121 方块 Q -> 牌对象（类型 1， 值 12）
// 支持单张或多张
type Card struct {
	Type    int64 // 牌的类型， 1：方块♦，2：梅花♣，3：红桃♥，4：黑桃♠，5：小王🃏，6：大王🃏
	Value   int64 // 牌的值 小王：160，大王：170
	IsLaizi bool  // 是否癞子牌
	IsUse   bool  // 是否已经使用，复合牌型需要判断牌是否被使用
}

// 牌组排序
func (p *Poker) SortCards(cards []*Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Value < cards[j].Value {
			return true
		} else if cards[i].Value == cards[j].Value {
			return cards[i].Type < cards[j].Type
		}
		return false
	})
}

// 牌组打乱
func (p *Poker) Disrupt(nums []*Card) {
	rand.Seed(time.Now().UnixNano())
	m := len(nums)
	for i := 0; i < m; i++ {
		j := rand.Intn(m-i) + i
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 从牌组中随机几张牌
func (p *Poker) RandomGet(nums []*Card, n int) []*Card {
	// 打乱牌组
	p.Disrupt(nums)
	// 取前n张
	if len(nums) < n {
		return nums
	}
	return nums[:n]
}

// 从牌组中随机几张不相同的牌
func (p *Poker) RandomGetDifferent(nums []*Card, n int) []*Card {
	// 打乱牌组
	p.Disrupt(nums)
	var result []*Card
	var isExists = make(map[int64]struct{})
	var ok bool
	for _, item := range nums {
		if _, ok = isExists[item.Value]; !ok {
			isExists[item.Value] = struct{}{}
			result = append(result, item)
			if len(result) == n {
				break
			}
		}
	}
	return result
}

// NumToCard 牌值转换为卡牌对象
// num: 可以是数组，也可以是单张牌的牌值
// return: 卡牌对象
func (p *Poker) NumToCard(num interface{}) []*Card {
	var cards []*Card
	switch num.(type) {
	case int64:
		v := num.(int64)
		card := Card{}
		card.Value = int64(v / 10)
		card.Type = v - card.Value*10
		if 160 == v {
			card.Type = 5
			card.Value = 16
		}
		if 170 == v {
			card.Type = 6
			card.Value = 17
		}
		if p.IsLaizi(card.Value) {
			card.IsLaizi = true
		}
		cards = append(cards, &card)
	case []int64:
		vs := num.([]int64)
		for _, v := range vs {
			card := p.NumToCard(v)
			cards = append(cards, card...)
		}
	default:
		return nil
	}
	return cards
}

// CardToNum 扑克牌对象转换为数字表示方式的值
// cards: 扑克牌对象。可以是单个卡牌对象，也可以是卡牌类型的切片 []*Card
// return: 扑克牌对象数字表示方式的值，因为可以有多个输入值，所以返回的是一个数组
func (p *Poker) CardToNum(cards interface{}) []int64 {
	var (
		numCards []int64
	)
	switch cards.(type) {
	case Card:
		card := cards.(Card)
		numCard := card.Type + card.Value*10
		if card.Type == 5 && card.Value == 16 {
			numCard = 160
		} else if card.Type == 6 && card.Value == 17 {
			numCard = 170
		}
		numCards = append(numCards, numCard)
	case *Card:
		card := cards.(*Card)
		numCard := card.Type + card.Value*10
		if card.Type == 5 && card.Value == 16 {
			numCard = 160
		} else if card.Type == 6 && card.Value == 17 {
			numCard = 170
		}
		numCards = append(numCards, numCard)
	case []*Card:
		cs := cards.([]*Card)
		for _, c := range cs {
			numCards = append(numCards, p.CardToNum(*c)...)
		}
	default:
	}
	return numCards
}

func (p *Poker) GetMaxNoJoker(cards []*Card) int64 {
	var result int64
	for _, item := range cards {
		if result < item.Value && item.Value < littleKing {
			result = item.Value
		}
	}
	return result
}

// CardBinarySearch cards的二分查找
func (p *Poker) CardBinarySearch(cards []*Card, target int64) *Card { // 二分查找
	low := 0
	high := len(cards) - 1

	for low <= high {
		mid := (low + high) / 2
		if cards[mid].Value == target {
			return cards[mid]
		} else if cards[mid].Value < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil // target not found
}

// GetCards 获取拥有的卡牌
func (p *Poker) GetCards(cards []*Card, target int64, num int) []*Card {
	if num <= 0 {
		num = len(cards)
	}
	var result []*Card
	for i := range cards {
		if cards[i].Value == target {
			result = append(result, cards[i])
			if len(result) == num {
				break
			}
		}
	}
	return result
}

// GetCards 获取n张癞子牌
func (p *Poker) GetCardsForLaizi(cards []*Card, target int64) []*Card {
	var result []*Card
	for i := range cards {
		if cards[i].IsLaizi {
			result = append(result, cards[i])
			if len(result) == int(target) {
				break
			}
		}
	}
	return result
}

// 清楚卡组使用状态
func (p *Poker) Use(cards []*Card, target []*Card) {
	for i := range cards {
		for i2 := range target {
			if cards[i].Type == target[i2].Type && cards[i].Value == target[i2].Value {
				cards[i].IsUse = true
			}
		}
	}
}

// 清楚卡组使用状态
func (p *Poker) UnUse(cards []*Card) {
	for i := range cards {
		cards[i].IsUse = false
	}
}

// 清楚卡组使用状态
func (p *Poker) IsUse(cards []*Card) bool {
	for i := range cards {
		if cards[i].IsUse {
			return true
		}
	}
	return false
}
