package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	clientCertFile = "client.crt"
	clientKeyFile  = "client.key"
	caFile         = "ca.crt" // CA-сертификат для проверки серверного сертификата
)

func client() {
	// 1. Загрузка сертификатов и ключей
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatalf("Ошибка загрузки клиентского сертификата: %v", err)
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
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool, // Проверка серверного сертификата
	}

	// 3. Установление TLS-соединения
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		log.Fatalf("Ошибка установления соединения: %v", err)
	}
	defer conn.Close()

	// 4. Обмен данными
	fmt.Fprintf(conn, "Hello, server!\n")
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalf("Ошибка чтения данных: %v", err)
	}

	fmt.Printf("Получено от сервера: %s\n", string(data))
}
