package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func encryptAES(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return hex.EncodeToString(cipherText), nil
}

func decryptAES(cipherTextHex, key string) (string, error) {
	cipherText, _ := hex.DecodeString(cipherTextHex)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func main() {
	var text, key string
	fmt.Print("Enter string to encrypt: ")
	fmt.Scan(&text)
	fmt.Print("Enter 16-byte key for encryption: ")
	fmt.Scan(&key)

	encrypted, err := encryptAES(text, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted:", encrypted)

	decrypted, err := decryptAES(encrypted, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", decrypted)
}
