package poker

import (
	"math/rand"
	"time"
)

func (p *Poker) ShuffleRandom() {
	p.lock.RLock()
	defer p.lock.RUnlock()
	var length = len(p.cards)
	for i := 0; i < length; i++ {
		j := rand.Intn(length)
		p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
	}
}

// 通过配置牌型出现概率来洗牌
// data:牌型出现的概率, 0:牌型，1：出现的概率，2：出现的数量
// userNum:游戏人数
// handNum:手牌数量
func (p *Poker) ShuffleProbability(data [][3]int64) {
	p.ShuffleRandom()
	if len(data) == 0 {
		return
	}
	p.lock.RLock()
	defer p.lock.RUnlock()
	var cardSet = make([][]*Card, 0)
	var r = rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	for i := range data {
		if data[i][0] == JokePair { // 出现双王
			if r.Float64() < float64(data[i][1])/100 { // 命中
				if pop, success := p.CardsAssignValuePop([]int64{littleKing, bigKing}); success {
					cardSet = append(cardSet, pop)
				}
			} else {
				continue
			}
		}
		if data[i][0] == Bomb { // 出现炸弹的概率
			var baseRand, tmpRand int64
			if r.Float64() < float64(data[i][1])/100 { // 命中
				for j := int64(0); j < data[i][2]; j++ {
					baseRand = r.Int63() % 15
					tmpRand = baseRand + 1
					for baseRand != tmpRand {
						if pop, success := p.CardsAssignValuePop([]int64{tmpRand, tmpRand, tmpRand, tmpRand}); success {
							cardSet = append(cardSet, pop)
							break
						} else {
							tmpRand++
							if tmpRand > 15 {
								tmpRand = 0
							}
						}
					}
				}
			} else {
				continue
			}
		}
		if data[i][0] == TrioStraight { // 出现飞机的概率
			var baseRand, tmpRand, tmpRand2 int64
			if r.Float64() < float64(data[i][1])/100 { // 命中
				for j := int64(0); j < data[i][2]; j++ {
					baseRand = r.Int63() % 15
					tmpRand = baseRand + 1
					for baseRand != tmpRand {
						tmpRand2 = tmpRand + 1
						if pop, success := p.CardsAssignValuePop([]int64{tmpRand, tmpRand, tmpRand, tmpRand2, tmpRand2, tmpRand2}); success {
							cardSet = append(cardSet, pop)
							break
						} else {
							tmpRand++
							if tmpRand > 15 {
								tmpRand = 0
							}
						}
					}
				}
			} else {
				continue
			}
		}
		if data[i][0] == Trio { // 出现三张的概率
			var baseRand, tmpRand int64
			if r.Float64() < float64(data[i][1])/100 { // 命中
				for j := int64(0); j < data[i][2]; j++ {
					baseRand = r.Int63() % 15
					tmpRand = baseRand + 1
					for baseRand != tmpRand {
						if pop, success := p.CardsAssignValuePop([]int64{tmpRand, tmpRand, tmpRand}); success {
							cardSet = append(cardSet, pop)
							break
						} else {
							tmpRand++
							if tmpRand > 15 {
								tmpRand = 0
							}
						}
					}
				}
			} else {
				continue
			}
		}
		if data[i][0] == SingleStraight { // 顺子
			var baseRand, tmpRand int64
			if r.Float64() < float64(data[i][1])/100 { // 命中
				for j := int64(0); j < data[i][2]; j++ {
					baseRand = r.Int63() % 15
					tmpRand = baseRand + 1
					for baseRand != tmpRand && baseRand >= there && baseRand+4 < two {
						if pop, success := p.CardsAssignValuePop([]int64{tmpRand, tmpRand + 1, tmpRand + 2, tmpRand + 3, tmpRand + 4}); success {
							cardSet = append(cardSet, pop)
							break
						} else {
							tmpRand++
							if tmpRand > 15 {
								tmpRand = 0
							}
						}
					}
				}
			} else {
				continue
			}
		}
		if data[i][0] == PairStraight { // 连对
			var baseRand, tmpRand int64
			if r.Float64() < float64(data[i][1])/100 { // 命中
				for j := int64(0); j < data[i][2]; j++ {
					baseRand = r.Int63() % 15
					tmpRand = baseRand + 1
					for baseRand != tmpRand && baseRand >= there && baseRand+2 < two {
						if pop, success := p.CardsAssignValuePop([]int64{tmpRand, tmpRand, tmpRand + 1, tmpRand + 1, tmpRand + 2, tmpRand + 2}); success {
							cardSet = append(cardSet, pop)
							break
						} else {
							tmpRand++
							if tmpRand > 15 {
								tmpRand = 0
							}
						}
					}
				}
			} else {
				continue
			}
		}
	}
	for i := range p.cards {
		cardSet = append(cardSet, []*Card{p.cards[i]})
	}
	// 打乱通过配置生出出来的牌
	rand.Shuffle(len(cardSet), func(i, j int) {
		cardSet[i], cardSet[j] = cardSet[j], cardSet[i]
	})
	var newCards []*Card
	for i := range cardSet {
		newCards = append(newCards, cardSet[i]...)
	}
	p.cards = newCards
}
