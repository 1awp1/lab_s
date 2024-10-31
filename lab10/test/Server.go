package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	certFile = "server.crt"
	keyFile  = "server.key"
)

func server() {
	// 1. Загрузка сертификатов и ключей
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Ошибка загрузки сертификата: %v", err)
	}

	// CA-сертификат для проверки серверного сертификата
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("Ошибка загрузки CA-сертификата: %v", err)
	}

	// Пул CA-сертификатов
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// 2. Создание конфигурации TLS
	serverConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // Взаимная аутентификация
		ClientCAs:    certPool,
	}

	// 3. Запуск сервера
	// Обработчик запросов
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n")
		fmt.Printf("Получен запрос от %s\n", r.RemoteAddr)
	})

	// Запуск TLS-сервера
	server := &http.Server{
		Addr:      ":8080",
		Handler:   handler,
		TLSConfig: serverConfig,
	}

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
