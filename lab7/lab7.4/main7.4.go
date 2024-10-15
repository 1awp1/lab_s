package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Привет, мир!")
}

func handleData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка считывания данных", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Получены данные:", data)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", handleHello)
	// curl -X POST http://localhost:8080/data -H "Content-Type: application/json" -d '{"key": "value"}'
	http.HandleFunc("/data", handleData)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
