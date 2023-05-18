package poker

import "sync"

const (
	there      = 3
	four       = 4
	five       = 5
	six        = 6
	seven      = 7
	eight      = 8
	nine       = 9
	ten        = 10
	jack       = 11
	queen      = 12
	king       = 13
	ace        = 14
	two        = 15
	littleKing = 16
	bigKing    = 17
)

type Poker struct {
	baseCards []*Card       // 原始牌组
	cards     []*Card       // 牌组
	laiNum    int64         // 癞子数量
	laizi     []int64       // 癞子牌
	lock      *sync.RWMutex // 锁
}

func NewPokerAlgorithm() *Poker {
	return &Poker{lock: new(sync.RWMutex)}
}

// 存储默认手牌
func (p *Poker) StorageBaseCards(baseCards []int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.baseCards = p.NumToCard(baseCards)
	p.cards = p.baseCards
}

func (p *Poker) CardsPop(num int64) []*Card {
	p.lock.Lock()
	defer p.lock.Unlock()
	if num > int64(len(p.cards)) {
		return nil
	}
	result := p.cards[:num]
	p.cards = p.cards[num:]
	return result
}

// pop指定类型
func (p *Poker) CardsAssignPop(num []int64) bool {
	if len(num) == 0 {
		return true
	}
	if len(p.cards) == 0 {
		return false
	}
	var newCards []*Card
	var isExist bool
	var count int
	for i := range p.cards {
		isExist = false
		for j := range num {
			if p.CardToNum(p.cards[i])[0] == num[j] {
				isExist = true
				count++
			}
		}
		if !isExist {
			newCards = append(newCards, p.cards[i])
		}
	}
	if count == len(num) { // 校验是否满足输出
		p.cards = newCards
		return true
	} else {
		return false
	}
}

// pop指定类型
func (p *Poker) CardsAssignValuePop(num []int64) ([]*Card, bool) {
	if len(num) == 0 {
		return nil, true
	}
	if len(p.cards) == 0 {
		return nil, false
	}
	var newCards []*Card
	var isExist bool
	var result []*Card
	for i := range p.cards {
		isExist = false
		for j := range num {
			if p.cards[i].Value == num[j] {
				isExist = true
				result = append(result, p.cards[i])
				break
			}
		}
		if !isExist {
			newCards = append(newCards, p.cards[i])
		}
	}
	if len(result) == len(num) { // 校验是否满足输出
		p.cards = newCards
		return result, true
	} else {
		return nil, false
	}
}

func (p *Poker) CardsGetByValue(v int64) []*Card {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if v == 0 {
		return nil
	}
	p.SortCards(p.cards)
	var result []*Card
	for i := range p.cards {
		if p.cards[i].Value == v {
			result = []*Card{p.cards[i]}
			for j := i + 1; j < len(p.cards); j++ {
				if p.cards[j].Value == v {
					result = append(result, p.cards[j])
				} else {
					return result
				}
			}
		}
	}
	return nil
}
