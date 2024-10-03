package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	// Создание каналов для передачи случайных чисел и сообщений о четности
	numberChannel := make(chan int)
	parityChannel := make(chan string)

	// Горутина для генерации случайных чисел
	go func() {
		for i := 0; i <= 10; i++ {
			randomNumber := rand.Intn(100) // Генерируем случайное число от 0 до 99
			numberChannel <- randomNumber  // Отправляем число в numberChannel
			time.Sleep(time.Second)        // Пауза перед генерацией следующего числа
		}
		close(numberChannel) // Закрываем канал после отправки всех чисел
	}()

	// Горутина для определения четности/нечетности чисел
	go func() {
		for number := range numberChannel { // Цикл считывания чисел из numberChannel
			if number%2 == 0 {
				parityChannel <- "Четное"

			} else {
				parityChannel <- "Нечетное"

			}
		}
		close(parityChannel) // Закрываем канал после обработки всех чисел
	}()

	// Главный поток: обработка данных из каналов
	for i := 0; i <= 10; i++ {
		select {
		case number := <-numberChannel: // Считывание числа из numberChannel
			fmt.Println("Сгенерировано число:", number)
		case parity := <-parityChannel: // Считывание сообщения о четности из parityChannel
			fmt.Println("Четность:", parity)

		}
	}
}
