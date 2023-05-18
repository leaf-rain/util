package poker

const (
	LastUnConf int64 = -100 // 上家出牌不符合（游戏bug）
	Less       int64 = -1   // 出牌比上家小
	UnConf     int64 = 0    // 不符合牌型规则
	Greater    int64 = 1    // 出牌比上家大

	// fix: 1:带有癞子牌，3:纯癞子，4：默认没有癞子牌
	FixHave int64 = 1
	FixAll  int64 = 3
	FixNo   int64 = 4
)

// GetCardsFeature 获取牌型特征值
// nCards: 计算牌型的卡牌数据
// comboType: 牌的类型
// return: 牌型特征值
func (p *Poker) GetCardsFeature(nCards []int64) int64 {
	var length = len(nCards)
	cards := p.NumToCard(nCards)
	var cardType, section, cardValue, fix int64
	switch length {
	case 1: // 一张牌出现的情况只有单牌
		return Single
	case 2: // 两张牌出现的情况： 王炸、对子
		if cardType, section, cardValue, fix = p.isJokePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isOnePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	case 3: // 三张牌出现的情况： 三条
		if cardType, section, cardValue, fix = p.isTrio(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	case 4: // 四张牌出现的情况： 炸弹、三带单
		if cardType, section, cardValue, fix = p.isJokePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isBomb(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioWithSingle(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	case 5: // 五张牌出现的情况： 炸弹、三带双、顺子
		if cardType, section, cardValue, fix = p.isBomb(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioWithPair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isSingleStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	case 6: // 六张牌出现的情况： 炸弹、四带2单、飞机不带、连队、顺子
		if cardType, section, cardValue, fix = p.isJokePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isBomb(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isFourWithTwoSingle(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isPairStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isSingleStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	case 8: // 六张牌出现的情况： 炸弹、四带2单、飞机不带、连队、顺子
		if cardType, section, cardValue, fix = p.isJokePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isBomb(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isFourWithTwoPair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isPairStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isSingleStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioStraightWithSingle(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	default:
		if cardType, section, cardValue, fix = p.isJokePair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isBomb(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioStraightWithPair(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioStraightWithSingle(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isTrioStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isPairStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
		if cardType, section, cardValue, fix = p.isSingleStraight(cards); cardType != 0 {
			return p.EncodeFeature(cardType, int(section), cardValue, fix)
		}
	}
	return 0
}

// EncodeFeature 计算特征值
// cardType:卡牌组合类型
// section: 节数
// cardValue: 牌值
// fix: 1:带有癞子牌，3:纯癞子，4：默认没有癞子牌
func (p *Poker) EncodeFeature(cardType int64, section int, cardValue int64, fix int64) int64 {
	if cardType == 0 || cardValue == 0 {
		return 0
	}
	fixValue := int64(4)
	if fix != 0 {
		fixValue = fix
	}
	return int64(cardType)*1000000 + int64(section)*10000 + cardValue*10 + fixValue
}

// DecodeFeature 解析特征值
// feature: 特征值
// return: 1: 卡牌组合类型；2: 节数；3: 牌值
func (p *Poker) DecodeFeature(feature int64) (int64, int, int64, int64) {
	if feature == 0 {
		return 0, 0, 0, 0
	}
	cardType := feature / 1000000
	feature = feature % 1000000
	section := int(feature / 10000)
	feature = feature % 10000
	cardValue := feature / 10
	fix := feature % 10
	return cardType, section, cardValue, fix
}

// CompareFeature 比较特征值
// curFeature: 要使用牌型组合的特征值
// lastFeature: 要对比的特征值
// return: 大 （可以出牌） = 1；小 （不能出牌）=-1；牌型不符 = 0
func (p *Poker) CompareFeature(curFeature, lastFeature int64) int64 {
	if curFeature <= 0 {
		return Less
	}
	if lastFeature <= 0 {
		return Greater
	}
	var result int64
	fType, fSection, fValue, fFix := p.DecodeFeature(curFeature)
	tType, tSection, tValue, tFix := p.DecodeFeature(lastFeature)
	if fType == JokePair {
		if tType == JokePair {
			if fSection > tSection {
				return Greater
			} else {
				return Less
			}
		}
		return Greater
	}

	// 两个数量不同的炸弹比较
	if fType == tType && fType == Bomb {
		if fFix > tFix && fSection == tSection {
			return Greater
		}
		if fSection > tSection {
			return Greater
		} else {
			return Less
		}
	}

	// 类型不同，出牌为炸弹
	if fType != tType && fType == Bomb {
		result = Less
		// 出牌的类型大于目标牌的类型
		if fType > tType {
			result = Greater
		}
		return result
	}

	// 组合类型相同，检查具体值
	if fType == tType && fSection == tSection {
		result = Less
		// 要出的牌比目标牌型大
		if fValue > tValue {
			result = Greater
		}
		return int64(result)
	}

	if fType != tType || fSection != tSection {
		return UnConf
	}
	return UnConf
}
