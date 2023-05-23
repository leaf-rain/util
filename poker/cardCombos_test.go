package poker

import "testing"

func Test_HintCardCombo(t *testing.T) {
	p.SetLaizi([]int64{15})
	var c1 = []int64{152, 31, 32, 33}
	feature := p.GetCardsFeature(c1, 0)
	t.Log(feature)
	var cards = []int64{51, 52, 53, 54}
	result := p.HintCardCombo(cards, feature)
	t.Log(result)
}
