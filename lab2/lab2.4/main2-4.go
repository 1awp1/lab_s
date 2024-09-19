package main

import "fmt"

func strlength(str string) int {
	return len(str) // возрат длины строки
}

func main() {
	//4 задание
	var word string
	fmt.Scan(&word)
	result := strlength(word)
	fmt.Println("Длина строки:", result)

}
