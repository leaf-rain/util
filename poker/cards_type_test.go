package poker

import "testing"

func TestPoker_isBomb(t *testing.T) {
	p.SetLaizi([]int64{3, 4})
	var nCards = []int64{
		31, 32, 33, 34,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isBomb(cards))
	a, b, c, d := p.isBomb(cards)
	t.Log(p.EncodeFeature(a, int(b), c, d))
	t.Log(p.CompareFeature(41040054, 41040033))
	nCards = []int64{
		31, 32, 33, 54,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isBomb(cards))
}

func TestPoker_isJokePair(t *testing.T) {
	var nCards = []int64{
		160, 170, 170,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isJokePair(cards))
	nCards = []int64{
		31, 160, 170,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isJokePair(cards))
}

func TestPoker_isOnePair(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isOnePair(cards))
	nCards = []int64{
		31, 160, 170,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isOnePair(cards))
	nCards = []int64{
		31, 42,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isOnePair(cards))
	nCards = []int64{
		51, 42,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isOnePair(cards))
}

func TestPoker_isTrio(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170, 170,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrio(cards))
	nCards = []int64{
		31, 31, 31,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrio(cards))
	nCards = []int64{
		31, 42, 32,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrio(cards))
	nCards = []int64{
		51, 42, 52,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrio(cards))
}

func TestPoker_isTrioWithSingle(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170, 170, 31,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithSingle(cards))
	nCards = []int64{
		31, 31, 31, 31,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithSingle(cards))
	nCards = []int64{
		31, 42, 32, 33,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithSingle(cards))
	nCards = []int64{
		51, 42, 52, 31,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithSingle(cards))
}

func TestPoker_isTrioWithPair(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170, 170, 31, 32,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithPair(cards))
	nCards = []int64{
		31, 31, 31, 31, 52,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithPair(cards))
	nCards = []int64{
		42, 42, 42, 33, 31,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithPair(cards))
	nCards = []int64{
		51, 42, 52, 31, 31,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioWithPair(cards))
}

func TestPoker_isFourWithTwoSingle(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170, 170, 170, 31, 32,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoSingle(cards))
	nCards = []int64{
		31, 31, 31, 31, 52, 52,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoSingle(cards))
	nCards = []int64{
		31, 42, 32, 33, 53, 55,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoSingle(cards))
	nCards = []int64{
		31, 42, 42, 41, 41, 32,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoSingle(cards))
}

func TestPoker_isFourWithTwoPair(t *testing.T) {
	p.SetLaizi([]int64{4})
	var nCards = []int64{
		170, 170, 170, 170, 31, 32, 51, 52,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoPair(cards))
	nCards = []int64{
		31, 31, 31, 31, 52, 52, 41, 42,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoPair(cards))
	nCards = []int64{
		31, 42, 32, 33, 53, 55, 41, 42,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoPair(cards))
	nCards = []int64{
		31, 42, 42, 31, 31, 31, 41, 42,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isFourWithTwoPair(cards))
}

func TestPoker_isSingleStraight(t *testing.T) {
	p.SetLaizi([]int64{11})
	var nCards = []int64{
		111, 111, 111, 111, 71, 42, 51, 62,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isSingleStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 101, 42, 51, 81,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isSingleStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 81, 91, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isSingleStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 81, 91, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isSingleStraight(cards))
}

func TestPoker_isPairStraight(t *testing.T) {
	p.SetLaizi([]int64{11})
	var nCards = []int64{
		111, 111, 111, 111, 61, 42, 51, 111,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isPairStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 111, 52, 51, 81,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isPairStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 121, 131, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isPairStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 81, 91, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isPairStraight(cards))
}
func TestPoker_isTrioStraight(t *testing.T) {
	p.SetLaizi([]int64{11})
	var nCards = []int64{
		111, 111, 111, 111, 61, 42, 51, 111, 111,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 111, 52, 51, 61, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 121, 131, 141, 151, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraight(cards))
	nCards = []int64{
		111, 111, 111, 111, 81, 91, 141, 151, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraight(cards))
}
func TestPoker_isTrioStraightWithSingle(t *testing.T) {
	p.SetLaizi([]int64{11})
	var nCards = []int64{
		111, 111, 111, 111, 61, 42, 111, 111,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithSingle(cards))
	nCards = []int64{
		111, 111, 111, 111, 111, 52, 51, 61,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithSingle(cards))
	nCards = []int64{
		111, 111, 111, 111, 121, 131, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithSingle(cards))
	nCards = []int64{
		111, 111, 111, 111, 81, 91, 141, 151,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithSingle(cards))
}

func TestPoker_isTrioStraightWithPair(t *testing.T) {
	p.SetLaizi([]int64{11})
	var nCards = []int64{
		111, 111, 111, 111, 61, 42, 111, 111, 111, 111,
	}
	var cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithPair(cards))
	nCards = []int64{
		111, 111, 111, 111, 111, 52, 51, 61, 111, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithPair(cards))
	nCards = []int64{
		111, 111, 111, 101, 121, 121, 141, 141, 111, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithPair(cards))
	nCards = []int64{
		111, 111, 111, 91, 121, 131, 141, 151, 111, 111,
	}
	cards = p.NumToCard(nCards)
	t.Log(p.isTrioStraightWithPair(cards))
}
