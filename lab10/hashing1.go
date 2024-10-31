package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)

func hashData(input string, algo string) string {
	var h hash.Hash

	switch algo {
	case "MD5":
		h = md5.New()
	case "SHA-256":
		h = sha256.New()
	case "SHA-512":
		h = sha512.New()
	default:
		fmt.Println("Unsupported algorithm")
		return ""
	}

	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

func verifyHash(input string, hashValue string, algo string) bool {
	calculatedHash := hashData(input, algo)
	return calculatedHash == hashValue
}

func main() {
	var input, algo, hashValue string
	var choice int

	fmt.Println("1. Hash Data")
	fmt.Println("2. Verify Hash")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter data to hash: ")
		fmt.Scan(&input)
		fmt.Print("Enter hash algorithm (MD5, SHA-256, SHA-512): ")
		fmt.Scan(&algo)

		hashedData := hashData(input, algo)
		fmt.Println("Hashed Data:", hashedData)
	case 2:
		fmt.Print("Enter original data: ")
		fmt.Scan(&input)
		fmt.Print("Enter hash algorithm (MD5, SHA-256, SHA-512): ")
		fmt.Scan(&algo)
		fmt.Print("Enter hash value to verify: ")
		fmt.Scan(&hashValue)

		if verifyHash(input, hashValue, algo) {
			fmt.Println("Hash verified successfully.")
		} else {
			fmt.Println("Hash does not match.")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
