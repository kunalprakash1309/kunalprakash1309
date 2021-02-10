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
func EncryptString(key, text string) string {
	keys := []byte(key)
	texts := []byte(text)

	block, err := aes.NewCipher(keys)
	if err != nil {
		panic(err)
	}

	b := base64.StdEncoding.EncodeToString(texts)
	cipherText := make([]byte, aes.BlockSize + len(b))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))
	// fmt.Println(string(cipherText))
	// s := fmt.Sprintf("%0x", string(cipherText))
	// fmt.Println(s)
	// // return s
	// // fmt.Println(cipherText)
	// fmt.Println(string(cipherText))
	return string(cipherText)
}

// DecryptString decrypts the encrypted string to original
func DecryptString(key, text string) string {
	keys := []byte(key)
	texts := []byte(text)
	if key == "" || text == "" {
		return ""
	}
	block, err := aes.NewCipher(keys)
	if err != nil {
		panic(err)
	}
	if len(texts) < aes.BlockSize {
		panic(err)
	}
	iv := texts[:aes.BlockSize]
	texts = texts[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(texts, texts)
	data, err := base64.StdEncoding.DecodeString(string(texts))
	if err != nil {
		panic(err)
	}
	return string(data)
}