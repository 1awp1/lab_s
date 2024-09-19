package main

import (
	"fmt"
	"time"
)

func main() {
	//1 задание
	// Получаем текущее время и дату
	DateTime := time.Now()

	// Форматируем дату и время
	formattedTime := DateTime.Format("2006-01-02 15:04:05")

	// Выводим результат
	fmt.Println("Текущее дата и время:", formattedTime)
}
