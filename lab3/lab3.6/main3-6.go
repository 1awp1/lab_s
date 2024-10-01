package main

import "fmt"

func main() {
	// Создаем срез из строк
	strings := []string{"hello", "world", "this", "is", "a", "long", "string"}

	// Находим самую длинную строку
	longestString := ""
	for _, str := range strings {
		if len(str) > len(longestString) {
			longestString = str
		}
	}

	// Выводим самую длинную строку
	fmt.Println("Самая длинная строка:", longestString)
}
