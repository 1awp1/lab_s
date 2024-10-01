// main.go

package main

import (
	"fmt"
	"for_study/mathutils" // импорт нашего пакета
)

func main() {
	number := 5
	factorial := mathutils.Factorial(number)
	fmt.Printf("Факториал %d равен %d\n", number, factorial)
}
