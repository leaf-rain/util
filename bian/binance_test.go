package bian

import (
	"context"
	"os"
	"testing"
)

var binanceSrv binance
var ctx = context.Background()

func TestMain(m *testing.M) {
	binanceSrv = NewBinance()
	ret := m.Run()
	os.Exit(ret)
}
