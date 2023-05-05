package poker

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
	baseCards []*Card // 原始牌组
	laiNum    int64   // 癞子数量
	laizi     []int64 // 癞子牌
}

func NewPokerAlgorithm() *Poker {
	return &Poker{}
}

// 存储默认手牌
func (p *Poker) StorageBaseCards(baseCards []int64) {
	p.baseCards = p.NumToCard(baseCards)
}
