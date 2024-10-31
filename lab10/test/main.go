package main

import (
	"fmt"
	"log"
	"os"
	// Импортируем функции из файла certgen.go
	// Импортируем функцию server() из файла server.go
	// Импортируем функцию client() из файла client.go
)

func main() {
	// 1. Генерация сертификатов (если они отсутствуют)
	if _, err := os.Stat("server.crt"); os.IsNotExist(err) {
		if err := generateSelfSignedCert("server.crt", "server.key"); err != nil {
			log.Fatalf("Ошибка генерации самозаверяющего сертификата: %v", err)
		}
	}

	if _, err := os.Stat("ca.crt"); os.IsNotExist(err) {
		if err := generateCA("ca.crt"); err != nil {
			log.Fatalf("Ошибка генерации CA-сертификата: %v", err)
		}
	}

	if _, err := os.Stat("client.crt"); os.IsNotExist(err) {
		if err := generateClientCert("client.crt", "client.key", "ca.crt"); err != nil {
			log.Fatalf("Ошибка генерации клиентского сертификата: %v", err)
		}
	}

	// 2. Выбор режима (сервер или клиент)
	fmt.Println("Выберите режим (server/client):")
	var mode string
	fmt.Scanln(&mode)

	switch mode {
	case "server":
		server()
	case "client":
		client()
	default:
		fmt.Println("Неверный режим.")
	}
}
