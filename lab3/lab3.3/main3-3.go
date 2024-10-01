// main.go

package main

import (
	"fmt"
	"for_study/stringutils" // импорт нашего пакета
)

func main() {
	inputString := "Hello, world!"
	reversedString := stringutils.Reverse(inputString)
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
