package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type AesKeyType int32

const (
	AES128 AesKeyType = iota
	AES192
	AES256
)

// Supported:
//   - AES128
//   - AES192
//   - AES256
func NewAesKey(t AesKeyType) (string, error) {
	var bLen int32
	switch t {
	case AES128:
		bLen = 16
	case AES192:
		bLen = 24
	case AES256:
		bLen = 32
	default:
		return "", errors.New("unsupported AES key length")
	}
	bStr := GenerateRandomString(int(bLen))

	return bStr, nil
}

func AesEncrypt(key, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	plainText, err = AesPKCS7Pad(plainText, blockSize)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, blockSize+sha256.Size+len(plainText))
	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	payload := cipherText[blockSize+sha256.Size:]
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(payload, plainText)

	hash := hmac.New(sha256.New, key)
	hash.Write(payload)
	copy(mac, hash.Sum(nil))

	return cipherText, nil
}

func AesDecrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	if len(cipherText) <= blockSize+sha256.Size {
		return nil, errors.New("ciphertext too short")
	}

	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	cipherText = cipherText[blockSize+sha256.Size:]

	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("ciphertext is not block-aligned, maybe corrupted")
	}

	hash := hmac.New(sha256.New, key)
	hash.Write(cipherText)
	if !hmac.Equal(hash.Sum(nil), mac) {
		return nil, errors.New("hmac failure, message corrupted")
	}

	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	plainText, err = AesPKCS7UnPad(plainText, blockSize)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func AesPKCS7Pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize >= 256 {
		return nil, fmt.Errorf("invalid block size: %d", blockSize)
	}

	paddingLen := blockSize - len(data)%blockSize

	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(data, padding...), nil
}

func AesPKCS7UnPad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 { // empty
		return nil, errors.New("unpad called on zero length byte array")
	}
	if length%blockSize != 0 {
		return nil, errors.New("data is not block-aligned")
	}

	paddingLen := int(data[length-1])
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	if paddingLen > blockSize || paddingLen == 0 || !bytes.HasSuffix(data, padding) {
		return nil, errors.New("invalid padding")
	}
	return data[:length-paddingLen], nil
}
