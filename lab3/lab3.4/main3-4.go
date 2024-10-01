package main

import "fmt"

func main() {
	// Создаем массив из 5 целых чисел
	numbers := make([]int, 5)

	// Заполняем массив значениями
	for i := 0; i < 5; i++ {
		numbers[i] = i * 10
	}

	// Выводим значения массива на экран
	fmt.Println("Массив:", numbers)
}
