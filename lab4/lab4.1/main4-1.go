package main

import "fmt"

func main() {
	// Создаем карту с именами и возрастами
	people := map[string]int{
		"Иван":  30,
		"Мария": 25,
		"Петр":  40,
	}

	// Добавляем нового человека
	people["Анна"] = 28

	// Выводим все записи
	fmt.Println("Список людей:")
	for name, age := range people {
		fmt.Printf("%s - %d лет\n", name, age)
	}
}
