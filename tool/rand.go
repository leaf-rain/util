package tool

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	//if min >= max || min == 0 || max == 0 {
	//	min = 1
	//	max = 100
	//}
	return rand.Int63n(max-min) + min
}
