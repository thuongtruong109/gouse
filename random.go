package gouse

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/google/uuid"
)

func RandNum(min, max int) int {
	// [min, max)
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0
	}

	return int(randomInt.Int64()) + min
}

func RandID() string {
	randomID := Sprint(time.Now().UnixNano())
	return randomID
}

func RandStr(n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += string(CHAIN_STR[RandNum(0, len(CHAIN_STR)-1)])
	}

	return result
}

func RandDigit(n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += string(CHAIN_NUM[RandNum(0, len(CHAIN_NUM)-1)])
	}

	return result
}

func UUID() string {
	uuid := uuid.New()
	return uuid.String()
}
