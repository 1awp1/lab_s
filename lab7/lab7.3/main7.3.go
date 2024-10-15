package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// Обработка сигнала прерывания (Ctrl+C) для graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Прием входящих соединений
	go func() {
		for {
			conn, err := listener.Accept() // Ожидание соединения
			if err != nil {
				fmt.Println("Ошибка приема соединения:", err)
				continue
			}

			// Обработка соединения в отдельной горутине
			go handleConnection(conn)
		}
	}()

	// Ожидание сигнала прерывания
	<-interrupt
	fmt.Println("Сервер останавливается...")

	// Прекращение приема новых соединений
	listener.Close()

	// Ждем завершения всех обработок соединений
	time.Sleep(time.Second)
	fmt.Println("Сервер остановлен.")
}
