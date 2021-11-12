package bian

import (
	"fmt"
	"github.com/leaf-rain/util/bian/dto"
	"testing"
)

func Test_binance_ListenKeyGenerate(t *testing.T) {
	result, _ := binanceSrv.ListenKeyGenerate(ctx)
	fmt.Println("result ======>", result)
}

func Test_binance_ListenKeyRenew(t *testing.T) {
	result := binanceSrv.ListenKeyRenew(ctx, &dto.ListenKeyValue{ListenKey: "OIXgtWVtkg3B1IqHSfbdfxdQj0BRK175pSWKnfrRupQBddzOlkQiF5MCg0ey"})
	fmt.Println("result ======>", result)
}

func Test_binance_ListenKeyDestroy(t *testing.T) {
	result := binanceSrv.ListenKeyDestroy(ctx, &dto.ListenKeyValue{ListenKey: "OIXgtWVtkg3B1IqHSfbdfxdQj0BRK175pSWKnfrRupQBddzOlkQiF5MCg0ey"})
	fmt.Println("result ======>", result)
}
