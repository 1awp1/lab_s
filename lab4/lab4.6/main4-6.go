package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbersStr := strings.Split(input, " ")
	var numbers []int

	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("Ошибка: Некорректный ввод:", err)
			return
		}
		numbers = append(numbers, number)
	}

	fmt.Println("Числа в обратном порядке:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
}
