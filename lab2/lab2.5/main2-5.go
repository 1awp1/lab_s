package main

import (
	"fmt"
)

// Структура Rectangle
type Rectangle struct {
	Width  int
	Height int
}

// Метод Area для вычисления площади
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	//5 задание

	var width, hight int
	// Создание экземпляра структуры Rectangle
	fmt.Scan(&width, &hight)
	rect := Rectangle{Width: width, Height: hight}

	// Вычисление площади и вывод результата
	area := rect.Area()
	fmt.Println("Площадь прямоугольника:", area)

}
