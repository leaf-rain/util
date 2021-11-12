package tool

import "math/big"

func getLen(str string) int64 {
	switch str {
	case "eth":
		return 18
	case "frt":
		return 18
	case "erc20_usdt":
		return 6
	case "trc20_usdt":
		return 6
	case "brc20_usdt":
		return 6
	case "btc":
		return 18
	case "rinkeby_eth":
		return 18
	case "bnb":
		return 18
	}
	return 0
}

func AmountMul1e12(amount string, str string) (string, bool) {
	num := getLen(str)
	if num == 0 {
		return "", false
	}
	a, ok := new(big.Int).SetString(amount, 0)
	if !ok {
		return "", false
	}
	if num == 6 {
		return amount, true
	} else {
		return new(big.Int).Mul(big.NewInt(1e+12), a).String(), true
	}
}

func AmountDiv1e12UInt64(amount string, str string) (uint64, bool) {
	num := getLen(str)
	if num == 0 {
		return 0, false
	}
	a, ok := new(big.Int).SetString(amount, 0)
	if !ok {
		return 0, false
	}
	if num > 6 {
		return new(big.Int).Div(a, big.NewInt(1e+12)).Uint64(), true
	} else {
		return a.Uint64(), true
	}
}
