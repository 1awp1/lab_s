package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для реверсирования строки
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Функция-воркер для обработки задач
func worker(id int, tasks <-chan string, results chan<- string) {
	for task := range tasks {
		// Обработка задачи (реверсирование строки)
		reversed := reverseString(task)
		fmt.Printf("Рабочий %d: Обработана строка '%s' -> '%s'\n", id, task, reversed)

		// Отправка результата в канал results
		results <- reversed
	}
}

func main() {
	// Чтение количества воркеров от пользователя
	fmt.Print("Введите количество воркеров: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numWorkers, _ := strconv.Atoi(strings.TrimSpace(input))

	// Создание каналов для задач и результатов
	tasks := make(chan string)
	results := make(chan string)

	// Создание пула воркеров
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, tasks, results) // Запуск каждого воркера в отдельной горутине
	}

	// Чтение строк из файла
	file, err := os.Open("project\\lab6\\lab6.6\\text.txt") // Замените "input.txt" на имя вашего файла
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- scanner.Text() // Отправка задачи в канал tasks
	}

	// Закрытие канала tasks для сигнализации окончания отправки задач
	close(tasks)

	// Сбор результатов
	for i := 0; i < numWorkers; i++ {
		fmt.Println("Результат", <-results) // Считываем результат из канала results
	}

	// Закрытие канала results
	close(results)

	fmt.Println("Все рабочие завершили обработку.")
}
