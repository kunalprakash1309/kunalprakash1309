package main

import (
	"log"

	"github.com/kunalprakash1309/Microservice/encryptString/utils"
)

// AES keys should be of length 16, 24, 32
func main() {
	key := "passphrasewhichneedstobe32bytes!"
	message := "Kunal Prakash"
	log.Println("Original message: ", string(message))
	encryptedString := utils.EncryptString(key, message)
	log.Printf("Encrypted message: %0x\n", []byte(encryptedString))
	decryptedString := utils.DecryptString(key, encryptedString)
	//fmt.Println(encryptedString)
    // decryptedString := utils.DecryptString(key, []byte(encryptedString))
	log.Println("Decrypted message: ", decryptedString)
}