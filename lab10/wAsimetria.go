package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, _ := os.Create(fileName)
	defer outFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	pem.Encode(outFile, block)
}

func savePublicPEMKey(fileName string, pubkey *rsa.PublicKey) {
	pubkeyBytes, _ := x509.MarshalPKIXPublicKey(pubkey)
	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubkeyBytes,
	}

	pubkeyFile, _ := os.Create(fileName)
	defer pubkeyFile.Close()

	pem.Encode(pubkeyFile, block)
}

func signMessage(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(message)
	digest := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
}

func verifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) error {
	hash := sha256.New()
	hash.Write(message)
	digest := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest, signature)
}

func main() {
	// Генерация ключей
	privateKey, publicKey := generateKeys()
	savePEMKey("private.pem", privateKey)
	savePublicPEMKey("public.pem", publicKey)

	// Ввод сообщения вручную
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a message to sign: ")
	message, _ := reader.ReadString('\n')
	message = message[:len(message)-1] // Убираем символ новой строки

	fmt.Println("Original message:", message)

	// Подпись сообщения
	signature, _ := signMessage(privateKey, []byte(message))
	fmt.Println("Signature:", signature)

	// Проверка подписи
	err := verifySignature(publicKey, []byte(message), signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
	} else {
		fmt.Println("Signature verified.")
	}
}
