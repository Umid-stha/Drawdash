package main

import (
	"math/rand/v2"
)

func generateCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	code := make([]byte, 6)

	for i := range code {
		code[i] = chars[rand.Int32N(int32(len(chars)))]
	}

	return string(code)
}
