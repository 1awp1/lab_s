// tls_client.go

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

const (
	certFile = "server.crt"
	caFile   = "ca.crt" // Для взаимной аутентификации
)

func main() {
	// Загрузка сертификата сервера
	log.Println("Загрузка сертификата сервера...")
	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	// Загрузка сертификата клиента и ключа
	log.Println("Загрузка сертификата клиента и ключа...")
	clientCert, err := tls.LoadX509KeyPair("ca.crt", "ca.key")
	if err != nil {
		log.Fatal(err)
	}

	// Создание конфигурации TLS
	log.Println("Создание конфигурации TLS...")
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
	}

	// Установка соединения
	log.Println("Установление соединения...")
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Проверка сертификата сервера
	log.Println("Проверка сертификата сервера...")
	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		log.Println("Ошибка: сертификат сервера не найден.")
		return
	}

	log.Println("Сертификаты сервера:")
	for i := range state.PeerCertificates {
		cert := state.PeerCertificates[i]
		log.Printf(" %s\n", cert.Subject.CommonName)
	}

	// Отправка данных
	log.Println("Отправка данных...")
	message := []byte("Привет от клиента!")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	// Получение ответа
	log.Println("Получение ответа...")
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Получено от сервера: %s\n", buffer[:n])
}
