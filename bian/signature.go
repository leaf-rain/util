package bian

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (b binance) SignGet(queryString, requestBody string) string {
	return hmacSha256(queryString+requestBody, b.SecretKey)
}

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
