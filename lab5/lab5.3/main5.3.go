package main

import (
	"fmt"
	"math"
)

// Структура Circle с полем radius
type Circle struct {
	radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Создание экземпляра структуры Circle
	var radius float64
	fmt.Scan(&radius)
	circle := Circle{radius: radius}

	// Вычисление площади круга
	area := circle.Area()

	// Вывод площади круга
	fmt.Printf("Площадь круга: %.2f\n", area)
}
