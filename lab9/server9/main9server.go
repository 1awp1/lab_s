package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

// Структура пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Структура для хранения пользователей
var users = make(map[int]User)
var userIDCounter = 1
var mu sync.Mutex

// Структура для аутентификации
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Структура для JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Получение всех пользователей
func getUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	usersList := make([]User, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usersList); err != nil {
		http.Error(w, "Unable to encode users", http.StatusInternalServerError)
		return
	}
}

// Добавление нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	mu.Lock()
	user.ID = userIDCounter
	users[userIDCounter] = user
	userIDCounter++
	mu.Unlock()

	json.NewEncoder(w).Encode(user)
}

// Обновление информации о пользователе
func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.ID]; exists {
		users[user.ID] = user
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}

}

// Удаление пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[user.ID]; exists {
		delete(users, user.ID)
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	// Декодируем входящие данные
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Проверяем учетные данные пользователя
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			// Создаем токен
			expirationTime := time.Now().Add(5 * time.Minute)
			claims := &Claims{
				Username: creds.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				http.Error(w, "Could not generate token", http.StatusInternalServerError)
				return
			}

			// Устанавливаем токен в куки
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			// Возвращаем информацию о пользователе
			response := User{ID: user.ID, Username: user.Username}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response) // Отправляем ID пользователя

			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "No token", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
func registerUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Проверяем, существует ли пользователь с таким именем
	for _, u := range users {
		if u.Username == user.Username {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}

	user.ID = userIDCounter
	users[userIDCounter] = user
	userIDCounter++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/register", registerUser) // Публичный маршрут

	// Авторизация и защищенные маршруты
	http.HandleFunc("/login", login)
	http.HandleFunc("/users", authenticate(getUsers))
	http.HandleFunc("/users/create", authenticate(createUser))
	http.HandleFunc("/users/update", authenticate(updateUser))
	http.HandleFunc("/users/delete", authenticate(deleteUser))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
