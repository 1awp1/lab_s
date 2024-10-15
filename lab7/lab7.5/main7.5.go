package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Method: %s, URL: %s, Duration: %s", r.Method, r.URL, duration)
	})
}

// Обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// Обработчик для страницы "О нас"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	mux := http.NewServeMux()

	// Регистрация маршрутов
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// Обертывание маршрутизатора в middleware
	loggedMux := loggingMiddleware(mux)

	// Запуск сервера
	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}
