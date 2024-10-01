package main

import "fmt"

// Структура Person с полями name и age
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) PrintInfo() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.name, p.age)
}

func main() {
	var name1 string
	var age1 int
	fmt.Scan(&name1, &age1)
	// Создание экземпляра структуры Person
	person := Person{name: name1, age: age1}

	// Вывод информации о человеке
	person.PrintInfo()
}
