package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/leaf-rain/util/tool"
)

func main() {
	stamp := tool.GetTimeUnixMilli()
	fmt.Println(stamp)
	//data := "coin=BTC&recvWindow=60000&timestamp="+tool.Int64ToStr(stamp)
	//data := "coin=USDT&limit=1000&status=1&recvWindow=60000&timestamp="+tool.Int64ToStr(stamp)
	data := "asset=USDT&needBtcValuation=true&recvWindow=60000&timestamp=" + tool.Int64ToStr(stamp)
	secret := ""
	fmt.Println(hmacSha256(data, secret))

}
func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
