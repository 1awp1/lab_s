package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для представления запроса
type Request struct {
	Operation string
	Num1      float64
	Num2      float64
}

// Структура для представления результата
type Result struct {
	Value     float64
	Error     error
	Operation string // Храним операцию для отладки
}

// Функция для обработки запросов на сложение
func add(req Request) Result {
	return Result{Value: req.Num1 + req.Num2, Operation: req.Operation}
}

// Функция для обработки запросов на вычитание
func subtract(req Request) Result {
	return Result{Value: req.Num1 - req.Num2, Operation: req.Operation}
}

// Функция для обработки запросов на умножение
func multiply(req Request) Result {
	return Result{Value: req.Num1 * req.Num2, Operation: req.Operation}
}

// Функция для обработки запросов на деление
func divide(req Request) Result {
	if req.Num2 == 0 {
		return Result{Error: fmt.Errorf("Деление на ноль"), Operation: req.Operation}
	}
	return Result{Value: req.Num1 / req.Num2, Operation: req.Operation}
}

// Функция-обработчик для калькулятора
func calculatorHandler(requests <-chan Request, results chan<- Result) {
	for req := range requests { // Цикл обработки запросов
		switch req.Operation {
		case "+":
			results <- add(req)
		case "-":
			results <- subtract(req)
		case "*":
			results <- multiply(req)
		case "/":
			results <- divide(req)
		default:
			results <- Result{Error: fmt.Errorf("Неверная операция: %s", req.Operation), Operation: req.Operation}
		}
	}
}

func main() {
	// Создание каналов для запросов и результатов
	requests := make(chan Request)
	results := make(chan Result)

	// Запуск обработчика калькулятора в отдельной горутине
	go calculatorHandler(requests, results)

	// Клиентская часть: отправка запросов
	scanner := bufio.NewScanner(os.Stdin) // Используем bufio.Scanner для ввода
	for {
		fmt.Print("Введите операцию (например, 10 + 5): ")
		scanner.Scan() // Читаем строку ввода

		input := scanner.Text() // Получаем строку ввода

		parts := strings.Split(input, " ") // Разделяем входные данные
		if len(parts) != 3 {
			fmt.Println("Некорректный формат ввода. Используйте: 'число операция число'")
			continue // Переход к следующему циклу
		}

		// Преобразование чисел в float64
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования первого числа:", err)
			continue
		}
		num2, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Ошибка преобразования второго числа:", err)
			continue
		}

		// Создание запроса и отправка в канал requests
		req := Request{Operation: parts[1], Num1: num1, Num2: num2}
		requests <- req

		// Получение результата из канала results
		res := <-results
		if res.Error != nil {
			fmt.Println("Ошибка:", res.Error)
		} else {
			fmt.Println("Результат:", res.Value)
		}
	}
}
