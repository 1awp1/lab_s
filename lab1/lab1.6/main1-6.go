package main

import (
	"fmt"
)

func main() {
	// 6 задание среднее значение
	var number1, number2, number3 int
	fmt.Println("Введите 3 числа подряд: ")
	fmt.Scan(&number1, &number2, &number3)
	sum := (number1 + number2 + number3) / 3
	fmt.Println("Сред арифмет:", sum)
}
