package main

import "fmt"

func main() {
	// Создаем массив из целых чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем срез из массива
	slice := numbers[1:3]
	fmt.Println("Срез:", slice) // Вывод: Срез: [2 3]

	// Добавление элементов в срез
	slice = append(slice, 6)
	fmt.Println("Срез после добавления:", slice) // Вывод: Срез после добавления: [2 3 6]

	// Удаление элемента из среза
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления:", slice) // Вывод: Срез после удаления: [2 6]
}
