package poker

import "testing"

func TestPoker_Shuffle(t *testing.T) {
	p.StorageBaseCards(Cards)
	p.ShuffleRandom()
	t.Log(p.CardToNum(p.cards))
}

func TestPoker_ShuffleProbability(t *testing.T) {
	p.StorageBaseCards(Cards)
	var req = [][3]int64{
		{JokePair, 100, 1},
		{Bomb, 100, 4},
		{Trio, 100, 4},
		{PairStraight, 100, 2},
	}
	p.ShuffleProbability(req)
	t.Log(len(p.cards))
	//t.Log(result)
	a := p.cards[:17]
	p.SortCards(a)
	t.Log(p.CardToNum(a))

	a = p.cards[17:34]
	p.SortCards(a)
	t.Log(p.CardToNum(a))

	a = p.cards[34:51]
	p.SortCards(a)
	t.Log(p.CardToNum(a))

	a = p.cards[51:]
	p.SortCards(a)
	t.Log(p.CardToNum(a))
}
