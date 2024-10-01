package main

import "fmt"

// Интерфейс Stringer
type Stringer interface {
	String() string
}

// Структура Book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Реализация метода String() для структуры Book
func (b Book) String() string {
	return fmt.Sprintf("Название: %s, Автор: %s, Год: %d", b.Title, b.Author, b.Year)
}

func main() {
	// Создание экземпляра структуры Book
	book := Book{Title: "Война и мир", Author: "Лев Толстой", Year: 1869}

	// Вывод информации о книге с использованием метода String()
	fmt.Println(book)
}
