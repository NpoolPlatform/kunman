//nolint:nolintlint,gosec,gomnd
package encrypt

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Upper(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{65, 90}
	result := make([]byte, 26)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + r.Intn(26))
	}
	return result
}

func Number(size int) []byte {
	if size <= 0 || size > 10 {
		size = 10
	}
	warehouse := []int{48, 57}
	result := make([]byte, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + r.Intn(9))
	}
	return result
}

func Lower(size int) []byte {
	if size <= 0 || size > 26 {
		size = 26
	}
	warehouse := []int{97, 122}
	result := make([]byte, 26)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result[i] = uint8(warehouse[0] + r.Intn(26))
	}
	return result
}

func Salt() string {
	var result string
	lowers := string(Lower(6))
	result += lowers
	numbers := string(Number(6))
	result += numbers
	uppers := string(Upper(6))
	result += uppers

	return result
}

func appendSalt(content, salt string) []byte {
	m5 := md5.New()
	m5.Write([]byte(content))
	m5.Write([]byte(salt))

	return m5.Sum(nil)
}

func EncryptWithSalt(content, salt string) (string, error) { //nolint
	salted := appendSalt(content, salt)

	hash, err := bcrypt.GenerateFromPassword(salted, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("encrypt salted error: %v", err)
	}

	return string(hash), nil
}

func VerifyWithSalt(content, target, salt string) error {
	salted := appendSalt(content, salt)

	err := bcrypt.CompareHashAndPassword([]byte(target), salted)
	if err != nil {
		return fmt.Errorf("mismatch content and target")
	}
	return nil
}
