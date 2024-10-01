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

// Функция для вывода площади каждого объекта в срезе
func PrintShapeAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}

func main() {
	// Создание экземпляров Rectangle и Circle
	rectangle := Rectangle{width: 5.0, height: 10.0}
	circle := Circle{radius: 5.0}

	// Создание среза интерфейсов Shape
	shapes := []Shape{rectangle, circle}

	// Вывод площади каждого объекта в срезе
	PrintShapeAreas(shapes)
}
