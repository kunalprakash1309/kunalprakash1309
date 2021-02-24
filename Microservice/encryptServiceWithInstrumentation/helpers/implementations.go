package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	// "crypto/rand"
	"encoding/base64"
	"errors"
)

// EncryptServiceInstance is the implementation of interface for micro service
type EncryptServiceInstance struct {}

var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

var errEmpty = errors.New("Secret Key or Text should not be empty")

// Encrypt encrypts the string with the given key
func (EncryptServiceInstance) Encrypt(_ context.Context,key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
	// s := fmt.Sprintf("Your input string: %s and Your given Key: %s" , text, key)
	// return s, nil
}

// Decrypt decrypts the encrypted string to original
func (EncryptServiceInstance) Decrypt(_ context.Context, key, text string) (string, error) {
	if key == "" || text == "" {
		return "", errEmpty
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil

	// s := fmt.Sprintf("Your output string: %s and Your given Key: %s" , text, key)
	// return s, nil
}