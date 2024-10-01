package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text)) // Удаляем пробелы и переводим в верхний регистр
	fmt.Println(text)
}
