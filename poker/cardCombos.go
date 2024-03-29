package poker

/*
牌值  ---  combo表示
3         03
4         04
5         05
6         06
7         07
8         08
9         09
10        10
J         11
Q         12
K         13
A         14
2         15
小王       16
大王       17
*/

// CardCombo 列举出可用的牌型组合
type CardCombo struct {
	Feature int64   // 特征
	Cards   []int64 // 具体使用的卡牌
}

// cardCombo 列举出可用的牌型组合
type cardCombo struct {
	Feature int64   // 特征
	Cards   []*Card // 具体使用的卡牌
}

// HintCardCombo 返回传入的牌中所有可用的牌型组合
// numCards: 传入的牌的值
// feature: 需要对比的牌的特征值。如果为 0 则表示所有的牌型都可以使用。
func (p *Poker) HintCardCombo(numCards []int64, feature int64) *CardCombo {
	var cards = p.NumToCard(numCards)
	p.SortCards(cards)
	if feature == 0 { // 没有比较牌，自己出牌，优先出数量比较大的牌
		// 飞机带对
		p.UnUse(cards)
		rs := p.GetMinTrioStraightWithPair(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 飞机带单张
		p.UnUse(cards)
		rs = p.GetMinTrioStraightWithSingle(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 飞机
		p.UnUse(cards)
		rs = p.GetMinTrioStraight(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 连对
		p.UnUse(cards)
		rs = p.GetMinPairStraight(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 顺子
		p.UnUse(cards)
		rs = p.GetMinSingleStraight(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 三带一
		p.UnUse(cards)
		rs = p.GetMinTrioWithSingle(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 三带对
		p.UnUse(cards)
		rs = p.GetMinTrioWithPair(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 三条
		p.UnUse(cards)
		rs = p.GetMinTrio(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 对子
		p.UnUse(cards)
		rs = p.GetMinOnePair(cards, feature, true, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}

		// 单张
		p.UnUse(cards)
		rs = p.GetMinSingle(cards, feature, false, true, true)
		if rs != nil {
			return &CardCombo{
				Feature: rs.Feature,
				Cards:   p.CardToNum(rs.Cards),
			}
		}
	} else {
		cType, _, _, _ := p.DecodeFeature(feature)
		var result *cardCombo
		p.UnUse(cards)
		switch {
		case cType == Single || feature == 0: // 单张
			result = p.GetMinSingle(cards, feature, true, true, true)
		case cType == OnePair: // 一对
			result = p.GetMinOnePair(cards, feature, true, true, true)
		case cType == Trio: // 三条
			result = p.GetMinTrio(cards, feature, true, true, true)
		case cType == TrioWithSingle: // 三带单
			result = p.GetMinTrioWithSingle(cards, feature, true, true, true)
		case cType == TrioWithPair: // 三带对
			result = p.GetMinTrioWithPair(cards, feature, true, true, true)
		case cType == FourWithTwoSingle: // 四带单
			result = p.GetMinFourWithTwoSingle(cards, feature, true, true, true)
		case cType == FourWithTwoPair: // 四带对
			result = p.GetMinFourWithTwoPair(cards, feature, true, true, true)
		case cType == SingleStraight: // 单顺
			result = p.GetMinSingleStraight(cards, feature, true, true, true)
		case cType == PairStraight: // 连对
			result = p.GetMinPairStraight(cards, feature, true, true, true)
		case cType == TrioStraight: // 飞机
			result = p.GetMinTrioStraight(cards, feature, true, true, true)
		case cType == TrioStraightWithSingle: // 飞机带单
			result = p.GetMinTrioStraightWithSingle(cards, feature, true, true, true)
		case cType == TrioStraightWithPair: // 飞机带对
			result = p.GetMinTrioStraightWithPair(cards, feature, true, true, true)
		}
		if result != nil {
			return &CardCombo{
				Feature: result.Feature,
				Cards:   p.CardToNum(result.Cards),
			}
		}
	}
	// 炸弹
	p.UnUse(cards)
	rs := p.GetMinBomb(cards, feature, true)
	if rs != nil {
		return &CardCombo{
			Feature: rs.Feature,
			Cards:   p.CardToNum(rs.Cards),
		}
	}
	return nil
}

// HaveJoker 拥有王炸
func (p *Poker) HaveJoker(cards []*Card) bool {
	p.SortCards(cards)
	var length = len(cards)
	if length < 2 || cards[length-1].Value < littleKing || cards[length-2].Value < littleKing {
		return false
	}
	return true
}

// GetMinLaizi 获取最小的癞子
func (p *Poker) GetMinLaizi(cards []*Card, num int64) []*Card {
	if num == 0 || len(cards) == 0 {
		return nil
	}
	p.SortCards(cards)
	var count int64
	var result []*Card
	for i := range cards {
		if cards[i].IsUse {
			continue
		}
		if cards[i].IsLaizi {
			cards[i].IsUse = true
			result = append(result, cards[i])
			count++
			if count == num {
				return result
			}
		}
	}
	return nil
}

// GetMinBomb 获取最小的炸弹
func (p *Poker) GetMinBomb(cards []*Card, feature int64, isJoker bool) *cardCombo {
	// 优先使用癞子炸弹
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	valueSetSort(vs) // 按照次数排序
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var _, tmpSection, _, _ = p.DecodeFeature(feature)

	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i].isLaizi || vs[i].times >= int64(tmpSection) || (vs[i].value >= littleKing && vs[i].times > 1) { // 不用癞子和不拆炸弹
			continue
		}
		var bombCards []*Card
		if vs[i].times+laiziCount >= int64(tmpSection) {
			bombCards = append(bombCards, p.GetCards(cards, vs[i].value, 0)...)
			bombCards = append(bombCards, p.GetCardsForLaizi(cards, int64(tmpSection)-vs[i].times)...)
			cardType, section, cardValue, fix = p.isBomb(bombCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				return &cardCombo{
					Feature: newFeature,
					Cards:   bombCards,
				}
			}
		}
	}
	// 使用硬炸
	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i].times < int64(tmpSection) || (vs[i].value >= littleKing && vs[i].times > 1) { // 不用癞子和不拆炸弹
			continue
		}
		if vs[i].times >= int64(tmpSection) {
			var outCards = p.GetCards(cards, vs[i].value, 0)
			cardType, section, cardValue, fix = p.isBomb(outCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				return &cardCombo{
					Feature: newFeature,
					Cards:   outCards,
				}
			}
		}
	}
	// 使用王炸
	if isJoker {
		var bombs = p.GetCards(cards, littleKing, 0)
		bombs = append(bombs, p.GetCards(cards, bigKing, 0)...)
		if len(bombs) >= 2 {
			cardType, section, cardValue, fix = p.isJokePair(bombs)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				return &cardCombo{
					Feature: newFeature,
					Cards:   bombs,
				}
			}
		}
	}
	return nil
}

// GetMinSingle 获取单张牌型，优先给出最多只有一张的牌
// 单张顺序：硬单张、拆一对的硬单张、拆三条的硬单张、硬炸，拆同一硬炸的硬单张、再硬炸，拆同一硬炸的硬单张，
func (p *Poker) GetMinSingle(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	var haveJoker = p.HaveJoker(cards)
	// 先找单牌
	for i := range vs {
		if vs[i].isLaizi || (haveJoker && vs[i].value >= littleKing) { // 不用癞子
			continue
		}
		if vs[i].times == 1 {
			var tmpCard = p.CardBinarySearch(cards, vs[i].value)
			if tmpCard.IsUse {
				continue
			}
			tmpCards = []*Card{tmpCard}
			cardType, section, cardValue, fix = p.isSingle(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					tmpCard.IsUse = true
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	valueSetSort(vs) // 按照次数排序
	// 拆牌
	if divide {
		for i := len(vs) - 1; i >= 0; i-- {
			if vs[i].isLaizi || vs[i].times >= 4 || (vs[i].value >= littleKing && vs[i].times > 1) || (haveJoker && vs[i].value >= littleKing) { // 不用癞子和不拆炸弹
				continue
			}
			var tmpCard = p.CardBinarySearch(cards, vs[i].value)
			if tmpCard.IsUse {
				continue
			}
			tmpCards = []*Card{tmpCard}
			cardType, section, cardValue, fix = p.isSingle(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					tmpCard.IsUse = true
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi && laiziCount > 0 {
		for i := len(vs) - 1; i >= 0; i-- {
			if vs[i].isLaizi || (haveJoker && vs[i].value >= littleKing) {
				var tmpCard = p.CardBinarySearch(cards, vs[i].value)
				if tmpCard.IsUse {
					continue
				}
				tmpCards = []*Card{tmpCard}
				cardType, section, cardValue, fix = p.isSingle(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						tmpCard.IsUse = true
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinOnePair 获取对子
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinOnePair(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	var haveJoker = p.HaveJoker(cards)
	// 先找不用癞子的对子
	for i := range vs {
		if vs[i].isLaizi || (haveJoker && vs[i].value >= littleKing) { // 不用癞子
			continue
		}
		if vs[i].times == 2 {
			tmpCards = p.GetCards(cards, vs[i].value, 0)
			if p.IsUse(tmpCards) {
				continue
			}
			cardType, section, cardValue, fix = p.isOnePair(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	valueSetSort(vs) // 按照次数排序
	// 拆牌
	if divide {
		for i := len(vs) - 1; i >= 0; i-- {
			if vs[i].isLaizi || vs[i].times >= 4 || (vs[i].value >= littleKing && vs[i].times > 1) || (haveJoker && vs[i].value >= littleKing) { // 不用癞子和不拆炸弹
				continue
			}
			if vs[i].times > 2 {
				tmpCards = p.GetCards(cards, vs[i].value, 0)
				if p.IsUse(tmpCards) {
					continue
				}
				cardType, section, cardValue, fix = p.isOnePair(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi && laiziCount > 0 {
		var result *cardCombo
		var needLaizi int64 = 3 // 给默认值
		for i := len(vs) - 1; i >= 0; i-- {
			if vs[i].isLaizi && vs[i].times > 2 || (haveJoker && vs[i].value >= littleKing) {
				continue
			}
			var tmpNeedLaizi = 2 - vs[i].times
			if tmpNeedLaizi >= needLaizi {
				continue
			}
			laiziCards := p.GetMinLaizi(cards, tmpNeedLaizi)
			tmpCards = append([]*Card{p.CardBinarySearch(cards, vs[i].value)}, laiziCards...)
			if len(tmpCards) != 2 || p.IsUse(tmpCards) {
				continue
			}
			cardType, section, cardValue, fix = p.isOnePair(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					needLaizi = tmpNeedLaizi
					result = &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
		if result != nil {
			return result
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrio 获取三张组合
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrio(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	var haveJoker = p.HaveJoker(cards)
	// 不用癞子
	for i := range vs {
		if vs[i].isLaizi || (haveJoker && vs[i].value >= littleKing) { // 不用癞子
			continue
		}
		if vs[i].times == 3 {
			tmpCards = p.GetCards(cards, vs[i].value, 0)
			if p.IsUse(tmpCards) {
				continue
			}
			cardType, section, cardValue, fix = p.isTrio(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	valueSetSort(vs) // 按照次数排序

	// divide 没有必要拆炸弹

	// 使用癞子
	if laizi && laiziCount > 0 {
		var result *cardCombo
		var needLaizi int64 = 4 // 给默认值
		for i := len(vs) - 1; i >= 0; i-- {
			if vs[i].isLaizi && vs[i].times > 3 || (haveJoker && vs[i].value >= littleKing) {
				continue
			}
			var tmpNeedLaizi = 3 - vs[i].times
			if tmpNeedLaizi >= needLaizi {
				continue
			}
			laiziCards := p.GetMinLaizi(cards, tmpNeedLaizi)
			tmpCards = append([]*Card{p.CardBinarySearch(cards, vs[i].value)}, laiziCards...)
			if len(tmpCards) != 3 || p.IsUse(tmpCards) {
				continue
			}
			cardType, section, cardValue, fix = p.isTrio(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					needLaizi = tmpNeedLaizi
					result = &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
		if result != nil {
			return result
		}
	}

	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrioWithSingle 获取三张带单组合
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrioWithSingle(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	// 分解特征值
	var _, _, baseValue, _ = p.DecodeFeature(feature)
	var tmpValue = p.EncodeFeature(Trio, 1, baseValue, FixNo)
	// 先找不用癞子的
	var trio = p.GetMinTrio(cards, tmpValue, false, false, false)
	var single = p.GetMinSingle(cards, 0, false, false, false)
	if trio != nil && single != nil {
		tmpCards = append(trio.Cards, single.Cards...)
		cardType, section, cardValue, fix = p.isTrioWithSingle(tmpCards)
		newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
		if p.CompareFeature(newFeature, feature) == Greater {
			for i1 := range tmpCards {
				tmpCards[i1].IsUse = true
			}
			return &cardCombo{
				Feature: newFeature,
				Cards:   tmpCards,
			}
		}
	}

	// divide 拆牌，但是不拆炸弹
	if divide {
		trio = p.GetMinTrio(cards, tmpValue, false, true, false)
		single = p.GetMinSingle(cards, 0, false, true, false)
		if trio != nil && single != nil {
			tmpCards = append(trio.Cards, single.Cards...)
			cardType, section, cardValue, fix = p.isTrioWithSingle(tmpCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				for i1 := range tmpCards {
					tmpCards[i1].IsUse = true
				}
				return &cardCombo{
					Feature: newFeature,
					Cards:   tmpCards,
				}
			}
		}
	}

	// 使用癞子
	if laizi {
		trio = p.GetMinTrio(cards, tmpValue, false, false, true)
		single = p.GetMinSingle(cards, 0, false, false, true)
		if trio != nil && single != nil {
			tmpCards = append(trio.Cards, single.Cards...)
			cardType, section, cardValue, fix = p.isTrioWithSingle(tmpCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				for i1 := range tmpCards {
					tmpCards[i1].IsUse = true
				}
				return &cardCombo{
					Feature: newFeature,
					Cards:   tmpCards,
				}
			}
		}
	}

	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrioWithPair 获取三带两张组合
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrioWithPair(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	// 分解特征值
	var _, _, baseValue, _ = p.DecodeFeature(feature)
	var tmpValue = p.EncodeFeature(Trio, 1, baseValue, FixNo)
	// 先找不用癞子的
	var trio = p.GetMinTrio(cards, tmpValue, false, false, false)
	var tmp = p.GetMinOnePair(cards, 0, false, false, false)
	if trio != nil && tmp != nil {
		tmpCards = append(trio.Cards, tmp.Cards...)
		cardType, section, cardValue, fix = p.isTrioWithPair(tmpCards)
		newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
		if p.CompareFeature(newFeature, feature) == Greater {
			for i1 := range tmpCards {
				tmpCards[i1].IsUse = true
			}
			return &cardCombo{
				Feature: newFeature,
				Cards:   tmpCards,
			}
		}
	}

	// divide 拆牌，但是不拆炸弹
	if divide {
		trio = p.GetMinTrio(cards, tmpValue, false, true, false)
		tmp = p.GetMinOnePair(cards, 0, false, true, false)
		if trio != nil && tmp != nil {
			tmpCards = append(trio.Cards, tmp.Cards...)
			cardType, section, cardValue, fix = p.isTrioWithPair(tmpCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				for i1 := range tmpCards {
					tmpCards[i1].IsUse = true
				}
				return &cardCombo{
					Feature: newFeature,
					Cards:   tmpCards,
				}
			}
		}
	}

	// 使用癞子
	if laizi {
		trio = p.GetMinTrio(cards, tmpValue, false, false, true)
		tmp = p.GetMinOnePair(cards, 0, false, false, true)
		if trio != nil && tmp != nil {
			tmpCards = append(trio.Cards, tmp.Cards...)
			cardType, section, cardValue, fix = p.isTrioWithPair(tmpCards)
			newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
			if p.CompareFeature(newFeature, feature) == Greater {
				for i1 := range tmpCards {
					tmpCards[i1].IsUse = true
				}
				return &cardCombo{
					Feature: newFeature,
					Cards:   tmpCards,
				}
			}
		}
	}

	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinSingleStraight 获取单顺
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinSingleStraight(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	if baseSection == 0 {
		baseSection = len(cards)
	}
	var flag bool
	// 不用癞子,不拆牌
	for i := len(vs) - 1; i >= 0; i-- {
		p.UnUse(cards)
		if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值,不能接2以上
			continue
		}
		flag = true
		tmpCards = nil
		if vs[i].times == 1 && i+baseSection < len(vs) {
			for i1 := i - 1; i1 > i-baseSection && i1 >= 1 && vs[i1].times == 1 && !vs[i1].isLaizi; i1-- {
				if vs[i1].value != vs[i1-1].value+1 || vs[i1].value >= two {
					flag = false
					break
				}
				tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 1)...)
			}
			if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 5) {
				if p.IsUse(tmpCards) {
					continue
				}
				cardType, section, cardValue, fix = p.isSingleStraight(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 拆牌但是不拆炸弹和不使用癞子
	if divide {
		for i := len(vs) - 1; i >= 0; i-- {
			p.UnUse(cards)
			if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值
				continue
			}
			flag = true
			tmpCards = nil
			if vs[i].times < 4 && i+baseSection < len(vs) {
				for i1 := i - 1; i1 > i-baseSection && i1 >= 1 && vs[i1].times < 4 && !vs[i1].isLaizi; i1-- {
					if vs[i1].value != vs[i1-1].value+1 || vs[i1].value >= two {
						flag = false
						break
					}
					tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 1)...)
				}
				if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 5) {
					if p.IsUse(tmpCards) {
						continue
					}
					cardType, section, cardValue, fix = p.isSingleStraight(tmpCards)
					if cardType != 0 {
						newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
						if p.CompareFeature(newFeature, feature) == Greater {
							for i1 := range tmpCards {
								tmpCards[i1].IsUse = true
							}
							return &cardCombo{
								Feature: newFeature,
								Cards:   tmpCards,
							}
						}
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		var tmpLaiziCount = laiziCount
		for i := len(vs) - 1; i >= 0; i-- {
			p.UnUse(cards)
			if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值
				continue
			}
			flag = true
			tmpCards = nil
			if vs[i].times < 4 {
				for i1 := i - 1; i1 > i-baseSection && i1 >= 1 && vs[i1].times < 4 && !vs[i1].isLaizi && tmpLaiziCount >= 0; i1-- {
					if i1 < len(vs) && vs[i1].value < two && vs[i1].value == vs[i1-1].value+1 {
						tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 1)...)
					} else {
						if tmpLaiziCount < 1 {
							break
						}
						tmpLaiziCount--
						i1++ // 再判断一次
					}
				}
				tmpCards = append(tmpCards, p.GetMinLaizi(cards, laiziCount-tmpLaiziCount)...)
				if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 5) {
					if p.IsUse(tmpCards) {
						continue
					}
					if laiziCount-tmpLaiziCount >= 4 { // 使用癞子足够当炸弹使用了
						continue
					}
					cardType, section, cardValue, fix = p.isSingleStraight(tmpCards)
					if cardType != 0 {
						newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
						if p.CompareFeature(newFeature, feature) == Greater {
							for i1 := range tmpCards {
								tmpCards[i1].IsUse = true
							}
							return &cardCombo{
								Feature: newFeature,
								Cards:   tmpCards,
							}
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinPairStraight 获取连对
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinPairStraight(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	if baseSection == 0 {
		baseSection = len(cards) / 2
	}
	var flag bool
	// 不用癞子,不拆牌
	for i := len(vs) - 1; i >= 0; i-- {
		p.UnUse(cards)
		if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值,不能接2以上
			continue
		}
		flag = true
		tmpCards = nil
		if vs[i].times == 2 && i+baseSection < len(vs) {
			for i1 := i - 1; i1 > i-baseSection && i1 >= 1 && vs[i1].times == 1 && !vs[i1].isLaizi; i1-- {
				if vs[i1].value != vs[i1-1].value+1 || vs[i1].value >= two {
					flag = false
					break
				}
				tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 2)...)
			}
			if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 6) {
				if p.IsUse(tmpCards) {
					continue
				}
				cardType, section, cardValue, fix = p.isPairStraight(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 拆牌但是不拆炸弹和不使用癞子
	if divide {
		for i := len(vs) - 1; i >= 0; i-- {
			p.UnUse(cards)
			if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值
				continue
			}
			flag = true
			tmpCards = nil
			if vs[i].times < 4 && i+baseSection < len(vs) {
				for i1 := i - 1; i1 > i-baseSection && i1 >= 1 && vs[i1].times < 4 && !vs[i1].isLaizi; i1-- {
					if vs[i1].value != vs[i1-1].value+1 || vs[i1].value >= two {
						flag = false
						break
					}
					tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 2)...)
				}
				if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 6) {
					if p.IsUse(tmpCards) {
						continue
					}
					cardType, section, cardValue, fix = p.isPairStraight(tmpCards)
					if cardType != 0 {
						newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
						if p.CompareFeature(newFeature, feature) == Greater {
							for i1 := range tmpCards {
								tmpCards[i1].IsUse = true
							}
							return &cardCombo{
								Feature: newFeature,
								Cards:   tmpCards,
							}
						}
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		var tmpLaiziCount = laiziCount
		for i := len(vs) - 1; i >= 0; i-- {
			p.UnUse(cards)
			if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值
				continue
			}
			flag = true
			tmpCards = nil
			if vs[i].times < 4 {
				for i1 := i - 1; i1 < i-baseSection && i1 >= 1 && vs[i1].times < 4 && !vs[i1].isLaizi && tmpLaiziCount >= 0; i1-- {
					if i1 < len(vs) && vs[i1].value < two && vs[i1].value == vs[i1-1].value+1 {
						if tmpLaiziCount < 2-vs[i1].times {
							break
						}
						if vs[i1].times < 2 {
							tmpLaiziCount -= 2 - vs[i1].times
						}
						tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 2)...)
					} else {
						if tmpLaiziCount < 2 {
							break
						}
						tmpLaiziCount -= 2
						i1++ // 再判断一次
					}
				}
				tmpCards = append(tmpCards, p.GetMinLaizi(cards, laiziCount-tmpLaiziCount)...)
				if (flag && len(tmpCards)/2 == baseSection) || (feature == 0 && len(tmpCards) >= 6) {
					if p.IsUse(tmpCards) {
						continue
					}
					if laiziCount-tmpLaiziCount >= 4 { // 使用癞子足够当炸弹使用了
						continue
					}
					cardType, section, cardValue, fix = p.isPairStraight(tmpCards)
					if cardType != 0 {
						newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
						if p.CompareFeature(newFeature, feature) == Greater {
							for i1 := range tmpCards {
								tmpCards[i1].IsUse = true
							}
							return &cardCombo{
								Feature: newFeature,
								Cards:   tmpCards,
							}
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrioStraight 飞机不带
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrioStraight(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	vs, laiziCount, _ := cardsToValueSetOnLaizi(cards)
	var cardType, section, cardValue, fix int64
	var newFeature int64
	var tmpCards []*Card
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	if baseSection == 0 {
		baseSection = len(cards) / 3
	}
	var flag bool
	// 不用癞子,不拆牌
	for i := len(vs) - 1; i >= 0; i-- {
		p.UnUse(cards)
		if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值,不能接2以上
			continue
		}
		flag = true
		tmpCards = nil
		if vs[i].times == 3 && i+baseSection < len(vs) {
			for i1 := i + 1; i1 < i+baseSection && i1 >= 1 && vs[i1].times == 1 && !vs[i1].isLaizi; i1++ {
				if vs[i1].value != vs[i1-1].value+1 || vs[i1].value >= two {
					flag = false
					break
				}
				tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 3)...)
			}
			if (flag && len(tmpCards) == baseSection) || (feature == 0 && len(tmpCards) >= 6) {
				if p.IsUse(tmpCards) {
					continue
				}
				cardType, section, cardValue, fix = p.isTrioStraight(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 三条的不需要拆牌，因为只能拆炸弹

	// 使用癞子
	if laizi {
		var tmpLaiziCount = laiziCount
		for i := len(vs) - 1; i >= 0; i-- {
			p.UnUse(cards)
			if vs[i].isLaizi || vs[i].value <= baseValue || vs[i].value >= two { // 不用癞子,牌值小于最小值
				continue
			}
			flag = true
			tmpCards = nil
			if vs[i].times < 4 {
				for i1 := i - 1; i1 < i-baseSection && i1 >= 1 && vs[i1].times < 4 && !vs[i1].isLaizi && tmpLaiziCount >= 0; i1-- {
					if i1 < len(vs) && vs[i1].value < two && vs[i1].value == vs[i1-1].value+1 {
						if tmpLaiziCount < 3-vs[i1].times {
							break
						}
						if vs[i1].times < 3 {
							tmpLaiziCount -= 3 - vs[i1].times
						}
						tmpCards = append(tmpCards, p.GetCards(cards, vs[i1].value, 3)...)
					} else {
						if tmpLaiziCount < 3 {
							break
						}
						tmpLaiziCount -= 3
						i1++ // 再判断一次
					}
				}
				tmpCards = append(tmpCards, p.GetMinLaizi(cards, laiziCount-tmpLaiziCount)...)
				if (flag && len(tmpCards)/3 == baseSection) || (feature == 0 && len(tmpCards) >= 6) {
					if p.IsUse(tmpCards) {
						continue
					}
					if laiziCount-tmpLaiziCount >= 4 { // 使用癞子足够当炸弹使用了
						continue
					}
					cardType, section, cardValue, fix = p.isTrioStraight(tmpCards)
					if cardType != 0 {
						newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
						if p.CompareFeature(newFeature, feature) == Greater {
							for i1 := range tmpCards {
								tmpCards[i1].IsUse = true
							}
							return &cardCombo{
								Feature: newFeature,
								Cards:   tmpCards,
							}
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrioStraightWithSingle 飞机带单牌
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrioStraightWithSingle(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, cardValue, fix int64
	var section int64
	var newFeature int64
	var tmpCards []*Card
	var flag bool
	var tmpSeciton int
	p.UnUse(cards)
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	var trioFeature = p.EncodeFeature(Trio, baseSection, baseValue, FixNo)
	// 不拆牌和不使用癞子
	var trio = p.GetMinTrio(cards, trioFeature, false, false, false)
	if trio != nil {
		tmpCards = append(tmpCards, trio.Cards...)
		p.Use(cards, tmpCards)
		_, tmpSeciton, _, _ = p.DecodeFeature(trioFeature)
		for i := 0; i < tmpSeciton; i++ {
			single := p.GetMinSingle(cards, 0, false, true, true)
			if single != nil {
				tmpCards = append(tmpCards, single.Cards...)
			}
		}
		if (flag && len(tmpCards)/4 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
			cardType, section, cardValue, fix = p.isTrioStraightWithSingle(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		p.UnUse(cards)
		trio = p.GetMinTrio(cards, trioFeature, false, false, true)
		if trio != nil {
			tmpCards = append(tmpCards, trio.Cards...)
			p.Use(cards, tmpCards)
			_, tmpSeciton, _, _ = p.DecodeFeature(trioFeature)
			for i := 0; i < tmpSeciton; i++ {
				single := p.GetMinSingle(cards, 0, false, true, true)
				if single != nil {
					tmpCards = append(tmpCards, single.Cards...)
				}
			}
			if (flag && len(tmpCards)/4 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
				cardType, section, cardValue, fix = p.isTrioStraightWithSingle(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinTrioStraightWithPair 飞机带对子
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinTrioStraightWithPair(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, cardValue, fix int64
	var section int64
	var newFeature int64
	var tmpCards []*Card
	var flag bool
	var tmpSeciton int
	p.UnUse(cards)
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	var trioFeature = p.EncodeFeature(Trio, baseSection, baseValue, FixNo)
	// 不拆牌和不使用癞子
	var trio = p.GetMinTrio(cards, trioFeature, false, false, false)
	if trio != nil {
		tmpCards = append(tmpCards, trio.Cards...)
		p.Use(cards, tmpCards)
		_, tmpSeciton, _, _ = p.DecodeFeature(trioFeature)
		for i := 0; i < tmpSeciton; i++ {
			single := p.GetMinOnePair(cards, 0, false, true, true)
			if single != nil {
				tmpCards = append(tmpCards, single.Cards...)
			}
		}
		if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
			cardType, section, cardValue, fix = p.isTrioStraightWithPair(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		p.UnUse(cards)
		trio = p.GetMinTrio(cards, trioFeature, false, false, true)
		if trio != nil {
			tmpCards = append(tmpCards, trio.Cards...)
			p.Use(cards, tmpCards)
			_, tmpSeciton, _, _ = p.DecodeFeature(trioFeature)
			for i := 0; i < tmpSeciton; i++ {
				single := p.GetMinOnePair(cards, 0, false, true, true)
				if single != nil {
					tmpCards = append(tmpCards, single.Cards...)
				}
			}
			if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
				cardType, section, cardValue, fix = p.isTrioStraightWithPair(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinFourWithTwoSingle 四带单
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinFourWithTwoSingle(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, cardValue, fix int64
	var section int64
	var newFeature int64
	var tmpCards []*Card
	var flag bool
	p.UnUse(cards)
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	var trioFeature = p.EncodeFeature(Trio, baseSection, baseValue, FixNo)
	// 不拆牌和不使用癞子
	var trio = p.GetMinBomb(cards, trioFeature, false)
	if trio != nil {
		tmpCards = append(tmpCards, trio.Cards...)
		p.Use(cards, tmpCards)
		for i := 0; i < 2; i++ {
			single := p.GetMinSingle(cards, 0, false, true, true)
			if single != nil {
				tmpCards = append(tmpCards, single.Cards...)
			}
		}
		if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
			cardType, section, cardValue, fix = p.isFourWithTwoSingle(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		p.UnUse(cards)
		trio = p.GetMinTrio(cards, trioFeature, false, false, true)
		if trio != nil {
			tmpCards = append(tmpCards, trio.Cards...)
			p.Use(cards, tmpCards)
			for i := 0; i < 2; i++ {
				single := p.GetMinSingle(cards, 0, false, true, true)
				if single != nil {
					tmpCards = append(tmpCards, single.Cards...)
				}
			}
			if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
				cardType, section, cardValue, fix = p.isFourWithTwoSingle(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}

// GetMinFourWithTwoPair 四带两对
// cards: 卡牌数据
// feature: 调用传入的特征值
// bomb:是否使用炸弹，divide:是否拆牌,laizi:是否使用癞子
func (p *Poker) GetMinFourWithTwoPair(cards []*Card, feature int64, bomb, divide, laizi bool) *cardCombo {
	var cardType, cardValue, fix int64
	var section int64
	var newFeature int64
	var tmpCards []*Card
	var flag bool
	p.UnUse(cards)
	// 分解特征值
	var _, baseSection, baseValue, _ = p.DecodeFeature(feature)
	var trioFeature = p.EncodeFeature(Trio, baseSection, baseValue, FixNo)
	// 不拆牌和不使用癞子
	var trio = p.GetMinBomb(cards, trioFeature, false)
	if trio != nil {
		tmpCards = append(tmpCards, trio.Cards...)
		p.Use(cards, tmpCards)
		for i := 0; i < 2; i++ {
			single := p.GetMinOnePair(cards, 0, false, true, true)
			if single != nil {
				tmpCards = append(tmpCards, single.Cards...)
			}
		}
		if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
			cardType, section, cardValue, fix = p.isFourWithTwoPair(tmpCards)
			if cardType != 0 {
				newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
				if p.CompareFeature(newFeature, feature) == Greater {
					for i1 := range tmpCards {
						tmpCards[i1].IsUse = true
					}
					return &cardCombo{
						Feature: newFeature,
						Cards:   tmpCards,
					}
				}
			}
		}
	}
	// 使用癞子
	if laizi {
		p.UnUse(cards)
		trio = p.GetMinTrio(cards, trioFeature, false, false, true)
		if trio != nil {
			tmpCards = append(tmpCards, trio.Cards...)
			p.Use(cards, tmpCards)
			for i := 0; i < 2; i++ {
				single := p.GetMinOnePair(cards, 0, false, true, true)
				if single != nil {
					tmpCards = append(tmpCards, single.Cards...)
				}
			}
			if (flag && len(tmpCards)/6 == baseSection) || (feature == 0 && len(tmpCards) >= 8) {
				cardType, section, cardValue, fix = p.isFourWithTwoPair(tmpCards)
				if cardType != 0 {
					newFeature = p.EncodeFeature(cardType, int(section), cardValue, fix)
					if p.CompareFeature(newFeature, feature) == Greater {
						for i1 := range tmpCards {
							tmpCards[i1].IsUse = true
						}
						return &cardCombo{
							Feature: newFeature,
							Cards:   tmpCards,
						}
					}
				}
			}
		}
	}
	// 使用炸弹
	if bomb {
		if result := p.GetMinBomb(cards, feature, true); result != nil {
			return result
		}
	}
	return nil
}
