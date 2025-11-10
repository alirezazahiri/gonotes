package idutils

import (
	"math/rand"
	"time"
)

func GenerateIDInt64() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Int63()
}
