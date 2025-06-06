package usercode

import (
	"math/rand"
	"time"
)

func generate(length int) string {
	number := []byte("0123456789")
	var result []byte
	// nolint
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, number[r.Intn(len(number))])
	}
	return string(result)
}
