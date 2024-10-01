package main

import "fmt"

func main() {
	// Создаем карту с именами и возрастами
	people := map[string]int{
		"Иван":  30,
		"Мария": 25,
		"Петр":  40,
	}

	// Запрашиваем имя человека для удаления
	var nameToDelete string
	fmt.Print("Введите имя человека для удаления: ")
	fmt.Scanln(&nameToDelete)

	// Удаляем запись из карты
	delete(people, nameToDelete)

	// Выводим обновленный список людей
	fmt.Println("Обновленный список людей:")
	for name, age := range people {
		fmt.Printf("%s - %d лет\n", name, age)
	}
}
