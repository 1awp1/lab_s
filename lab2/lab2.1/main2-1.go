package main

import "fmt"

func main() {

	// 1 задание
	fmt.Println("Задание 1:")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}

}
