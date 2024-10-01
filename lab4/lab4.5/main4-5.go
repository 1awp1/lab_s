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
	input = strings.TrimSpace(input) // Удаляем пробелы в начале и конце строки

	numbers := strings.Split(input, " ") // Разделяем строку на числа
	var sum int

	for _, numberStr := range numbers {
		number, err := strconv.Atoi(numberStr) // Преобразуем строку в число
		if err != nil {
			fmt.Println("Ошибка: Некорректный ввод:", err)
			return
		}
		sum += number
	}

	fmt.Println("Сумма чисел:", sum)
}
