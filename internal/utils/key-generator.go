package utils

import "math/rand"

const keyLength = 12

const symbolsBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const symbolsCount = len(symbolsBytes)

func GenerateRandomKey() string {
	return randStringBytes(keyLength)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = symbolsBytes[rand.Intn(symbolsCount)]
	}
	return string(b)
}
