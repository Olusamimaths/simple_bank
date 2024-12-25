package util

import (
	"math/rand/v2"
	"strings"
)


func RandomInt(min, max int64) int64 {
	return rand.Int64N(max - min + 1) + min
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	length := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.IntN(length)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string {"NGN", "JPY", "USD", "CAD", "EUR"}
	length := len(currencies)
	
	return currencies[rand.IntN(length)]
}
