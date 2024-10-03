package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generateRandomNumbers(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100) // Generate random numbers between 0 and 99
	}
	return numbers
}

func sumSeries(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализирует генератор случайных чисел.

	// Start goroutines for each function
	go func() {
		fmt.Println("Calculating factorial...")
		time.Sleep(2 * time.Second) // Simulate delay
		result := factorial(5)
		fmt.Println("Factorial of 5:", result)
	}()

	go func() {
		fmt.Println("Generating random numbers...")
		time.Sleep(1 * time.Second) // Simulate delay
		numbers := generateRandomNumbers(5)
		fmt.Println("Random numbers:", numbers)
	}()

	go func() {
		fmt.Println("Calculating sum of series...")
		time.Sleep(3 * time.Second) // Simulate delay
		sum := sumSeries(10)
		fmt.Println("Sum of series:", sum)
	}()

	// Main goroutine waits for a short time before exiting
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine exiting...")
}
