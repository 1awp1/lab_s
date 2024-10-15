package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Адрес и порт сервера
	address := "localhost:8080" // Измените адрес и порт по необходимости

	// Подключение к серверу
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close() // Закрытие соединения после завершения работы

	// Создание буфера для ввода/вывода
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	// Цикл отправки сообщений и получения ответов
	for {
		fmt.Print("Введите сообщение: ")
		message, _ := reader.ReadString('\n')

		// Отправка сообщения на сервер
		_, err := writer.WriteString(message)
		if err != nil {
			fmt.Println("Ошибка отправки сообщения:", err)
			return
		}
		writer.Flush()

		// Получение ответа от сервера
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка получения ответа:", err)
			return
		}

		fmt.Println("Получен ответ:", response)

		// Выход из цикла после получения ответа
		break
	}
}
