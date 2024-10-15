package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // Закрытие соединения после обработки

	// Создание буфера для считывания сообщений
	reader := bufio.NewReader(conn)

	for {
		// Считывание сообщения от клиента с обработкой ошибки
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Соединение закрыто клиентом")
				return
			}
			fmt.Println("Ошибка чтения сообщения:", err)
			return
		}

		// Вывод сообщения на экран
		fmt.Println("Получено сообщение:", message)

		// Отправка подтверждения клиенту
		_, err = conn.Write([]byte("Сообщение получено!\n"))
		if err != nil {
			fmt.Println("Ошибка отправки подтверждения:", err)
		}
	}
}

func main() {
	// Задание адреса и порта для прослушивания
	address := "localhost:8080" // Измените адрес и порт по необходимости

	// Создание слушателя
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Ошибка создания слушателя:", err)
		return
	}
	defer listener.Close() // Закрытие слушателя после завершения работы

	fmt.Println("Сервер запущен на", address)

	// Прием входящих соединений
	for {
		conn, err := listener.Accept() // Ожидание соединения
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		// Обработка соединения в отдельной горутине
		go handleConnection(conn)
	}
}
