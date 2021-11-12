package bian

import (
	"fmt"
	"github.com/leaf-rain/util/bian/dto"
	"testing"
)

func Test_binance_RecordChargeGet(t *testing.T) {
	result, _ := binanceSrv.RecordChargeGet(ctx, &dto.RecordChargeGetReq{
		Coin:      "USDT",
		Status:    1,
		StartTime: 1631354565000,
		Limit:     1000,
		Offset:    0,
	})
	for _, item := range result {
		fmt.Println(fmt.Printf("resut ======> %+v", item))
	}
}

func Test_binance_RecordCashGet(t *testing.T) {
	result, _ := binanceSrv.RecordCashGet(ctx, &dto.RecordCashGetReq{
		Coin:      "USDT",
		Status:    1,
		StartTime: 1631354565000,
		Limit:     1000,
		Offset:    0,
	})
	for _, item := range result {
		fmt.Println(fmt.Printf("resut ======> %+v", item))
	}
}
