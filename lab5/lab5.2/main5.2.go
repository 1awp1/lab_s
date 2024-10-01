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

// Метод birthday для увеличения возраста на 1 год
func (p *Person) Birthday() {
	p.age++
}

func main() {
	var name1 string
	var age1 int
	fmt.Scan(&name1, &age1)
	// Создание экземпляра структуры Person
	person := Person{name: name1, age: age1}
	// Вывод информации о человеке до дня рождения
	fmt.Println("Информация до дня рождения:")
	person.PrintInfo()

	// Вызов метода birthday
	person.Birthday()

	// Вывод информации о человеке после дня рождения
	fmt.Println("Информация после дня рождения:")
	person.PrintInfo()
}
