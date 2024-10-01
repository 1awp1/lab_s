package main

import "fmt"

func averageAge(people map[string]int) float64 {
	var totalAge int
	var count int

	for _, age := range people {
		totalAge += age
		count++
	}

	if count == 0 {
		return 0 // Обработка случая, когда карта пуста
	}

	return float64(totalAge) / float64(count)
}

func main() {
	people := map[string]int{
		"Иван":  30,
		"Мария": 25,
		"Петр":  40,
	}

	average := averageAge(people)
	fmt.Println("Средний возраст:", average)
}
