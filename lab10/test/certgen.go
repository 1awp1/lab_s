package main

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func generateSelfSignedCert(certFile, keyFile string) error {
	// Создаем самозаверяющий сертификат
	cert, err := tls.X509KeyPair([]byte(certFile), []byte(keyFile))
	if err != nil {
		return err
	}

	// Создаем пул CA-сертификатов
	certPool := x509.NewCertPool()
	certPool.AddCert(cert.Leaf)

	// Создаем конфигурацию TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    certPool,
	}

	// Создаем сервер
	server := &http.Server{
		Addr:      ":8080",
		Handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello, world!\n") }),
		TLSConfig: tlsConfig,
	}

	// Запускаем сервер для генерации сертификата
	log.Fatal(server.ListenAndServeTLS("", ""))
	return nil
}

func generateCA(caFile string) error {
	// Создаем сертификат для CA
	caCert, err := tls.X509KeyPair(caFile, caFile)
	if err != nil {
		return err
	}

	// Создаем пул CA-сертификатов
	certPool := x509.NewCertPool()
	certPool.AddCert(caCert.Leaf)

	// Создаем конфигурацию TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{caCert},
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    certPool,
	}

	// Создаем сервер
	server := &http.Server{
		Addr:      ":8080",
		Handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello, world!\n") }),
		TLSConfig: tlsConfig,
	}

	// Запускаем сервер для генерации сертификата
	log.Fatal(server.ListenAndServeTLS("", ""))
	return nil
}

// Генерация сертификата для клиента
func generateClientCert(clientCertFile, clientKeyFile string, caCertFile string) error {
	// Загружаем CA-сертификат
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return err
	}

	// Создаем пул CA-сертификатов
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// Создаем TLS-конфигурацию
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	// Устанавливаем TLS-соединение с CA
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Создаем CSR (запрос на подпись сертификата)
	csr, err := generateCSR(clientCertFile, clientKeyFile)
	if err != nil {
		return err
	}

	// Отправляем CSR CA
	_, err = conn.Write([]byte(csr))
	if err != nil {
		return err
	}

	// Получаем подписанный сертификат от CA
	cert, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}

	// Сохраняем сертификат в файл
	err = ioutil.WriteFile(clientCertFile, cert, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Генерация CSR (запрос на подпись сертификата)
func generateCSR(certFile, keyFile string) (string, error) {
	// Загружаем ключ
	key, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return "", err
	}

	// Создаем CSR
	csr, err := x509.CreateCertificateRequest(x509.Certificate{
		Subject: pkix.Name{
			CommonName:   "Client",
			Organization: []string{"Example Inc."},
		},
	}, &key.PublicKey, &key.PrivateKey, nil)
	if err != nil {
		return "", err
	}

	// Кодируем CSR в PEM-формат
	csrPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	})

	return string(csrPem), nil
}
