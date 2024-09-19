package main

import "fmt"

// Получить среднее двух целых
func getSecondNumber(a int, b int) float64 {
	fmt.Scan(&a)
	fmt.Scan(&b)
	sum := (float64(a) + float64(b)) / 2
	return sum
}

func main() {

	// 6 задание
	var x, y int
	fmt.Println("Среднее двух целых", getSecondNumber(x, y))
}
