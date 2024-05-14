package utils

import "github.com/google/uuid"

const keyLength = 12

func GenerateRandomKey() string {
	return uuid.NewString()[:keyLength]
}
