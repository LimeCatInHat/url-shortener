package utils

import (
	"hash/fnv"
	"strconv"
)

func hash(value []byte) uint64 {
	h := fnv.New64a()
	h.Write(value)
	return h.Sum64()
}

func GenerateKey(value []byte) string {
	return strconv.FormatUint(hash(value), 16)
}
