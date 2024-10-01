package main

import (
	"fmt"
	"math"
)

// Интерфейс Shape с методом Area()
type Shape interface {
	Area() float64
}

// Структура Rectangle с полями width и height
type Rectangle struct {
	width  float64
	height float64
}

// Структура Circle с полем radius
type Circle struct {
	radius float64
}

// Реализация метода Area() для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Реализация метода Area() для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Создание экземпляров Rectangle и Circle
	rectangle := Rectangle{width: 5.0, height: 10.0}
	circle := Circle{radius: 5.0}

	// Вывод площади фигур
	fmt.Printf("Площадь прямоугольника: %.2f\n", rectangle.Area())
	fmt.Printf("Площадь круга: %.2f\n", circle.Area())

	// Использование интерфейса Shape для вычисления площади
	var shape Shape
	shape = rectangle
	fmt.Printf("Площадь фигуры (прямоугольник): %.2f\n", shape.Area())
	shape = circle
	fmt.Printf("Площадь фигуры (круг): %.2f\n", shape.Area())
}
