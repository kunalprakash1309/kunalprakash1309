package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Implements AES encryption algorithm

// var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// EncryptString encrypts the string with the given key
func EncryptString(key, text []byte) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	b := base64.StdEncoding.EncodeToString(text)
	cipherText := make([]byte, aes.BlockSize + len(b))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))
	return string(cipherText)

	// plainText := []byte(text)
	// cfb := cipher.NewCFBEncrypter(block, initVector)
	// cipherText := make([]byte, len(plainText))
	// cfb.XORKeyStream(cipherText, plainText)
	// return base64.StdEncoding.EncodeToString(cipherText)
}

// DecryptString decrypts the encrypted string to original
func DecryptString(key, text []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic(err)
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		panic(err)
	}
	return string(data)
	// block, err := aes.NewCipher([]byte(key))
	// if err != nil {
	// 	panic(err)
	// }

	// cipherText, _ := base64.StdEncoding.DecodeString(text)
	// cfb := cipher.NewCFBEncrypter(block, initVector)
	// plainText := make([]byte, len(cipherText))
	// cfb.XORKeyStream(plainText, cipherText)
	// return string(plainText)
}