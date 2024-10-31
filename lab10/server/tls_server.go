// tls_server.go

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const (
	certFile = "server.crt"
	keyFile  = "server.key"
	caFile   = "ca.crt" // Для взаимной аутентификации
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Использование: ./tls_server <порт>")
		os.Exit(1)
	}

	port := os.Args[1]

	// Загрузка сертификатов
	log.Println("Загрузка сертификатов...")
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Создание пула сертификатов для проверки сертификатов клиента
	log.Println("Создание пула сертификатов для проверки сертификатов клиента...")
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Создание конфигурации TLS
	log.Println("Создание конфигурации TLS...")
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // Взаимная аутентификация
		ClientCAs:    caCertPool,                     // Используйте caCertPool для проверки
	}

	// Создание сервера
	log.Println("Создание сервера...")
	listener, err := tls.Listen("tcp", "localhost:"+port, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("TLS-сервер запущен на порту %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка при приеме соединения:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Проверка сертификата клиента
	log.Println("Проверка сертификата клиента...")
	state := conn.(*tls.Conn).ConnectionState()
	if len(state.PeerCertificates) == 0 {
		log.Println("Ошибка: сертификат клиента не найден.")
		return
	}

	log.Println("Сертификаты клиента:")
	for i := range state.PeerCertificates {
		cert := state.PeerCertificates[i]
		log.Printf(" %s\n", cert.Subject.CommonName)
	}

	// Обработка данных
	log.Println("Обработка данных...")
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Ошибка чтения:", err)
			return
		}

		log.Printf("Получено от клиента: %s\n", buffer[:n])

		// Отправка ответа
		log.Println("Отправка ответа...")
		message := []byte("Привет от сервера!")
		_, err = conn.Write(message)
		if err != nil {
			log.Println("Ошибка записи:", err)
			return
		}
	}
}
