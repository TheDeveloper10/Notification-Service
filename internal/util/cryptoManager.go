package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"notification-service/internal/helper"
)

type CryptoManager struct {}

func (cm *CryptoManager) Encrypt(text *string, key *string) *string {
	bytes := []byte(*text)

	block, err := aes.NewCipher([]byte(*key))
	if helper.IsError(err) {
		return nil
	}

	cipherText := make([]byte, aes.BlockSize+len(bytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); helper.IsError(err) {
		return nil
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], bytes)

	res := base64.URLEncoding.EncodeToString(cipherText)
	return &res
}

func (cm *CryptoManager) Decrypt(cryptoText *string, key *string) *string {
	cipherText, _ := base64.URLEncoding.DecodeString(*cryptoText)

	block, err := aes.NewCipher([]byte(*key))
	if helper.IsError(err) {
		return nil
	}

	if len(cipherText) < aes.BlockSize {
		return nil
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	res := string(cipherText)
	return &res
}