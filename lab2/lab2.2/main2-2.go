package main

import "fmt"

func checkNumber(num int) string { // получение числа
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func main() {
	var num int
	//2 задание
	fmt.Println("Задание 2:")
	fmt.Scan(&num)
	fmt.Println(checkNumber(num))
}
