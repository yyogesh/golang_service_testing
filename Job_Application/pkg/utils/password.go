package utils

import (
	"math/rand"
	"strings"
)

func GenerateFromPassword(charCount int) string {
	const digits = "0123456789abcdef"

	var password strings.Builder
	password.Grow(charCount)

	for i := 0; i < charCount; i++ {
		password.WriteByte(digits[rand.Intn(len(digits))])
	}

	return password.String()
}
