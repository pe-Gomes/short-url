package util

import (
	"math/rand"
	"strings"
)

const chars = "abcdefghijklmnopqrstuvwxy1234567890"

func RandomString(n int) string {
	var stringBuilder strings.Builder
	k := len(chars)

	for i := 0; i < n; i++ {
		char := chars[rand.Intn(k)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}
