package poker

import (
	"fmt"
	"testing"
	"time"
)

var p = NewPokerAlgorithm()

func TestMain(m *testing.M) {
	var now = time.Now()
	m.Run()
	fmt.Println("执行耗时:", time.Since(now))
}

func Test_Poker(t *testing.T) {
	p.StorageBaseCards(Cards)
	result := p.HintCardCombo(Cards, 41040034)
	t.Log(result)
}
