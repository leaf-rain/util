package tool

import (
	"fmt"
	"testing"
)

func TestAmountMul1e12(t *testing.T) {
	fmt.Println(AmountDiv1e12UInt64("1000000000000000", "frt"))
	fmt.Println(AmountDiv1e12UInt64("10", "erc20_usdt"))
}
