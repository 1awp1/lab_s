package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var client = &http.Client{}
var sessionToken string
var currentUserID int // для хранения ID текущего пользователя после логина

func main() {
	var choice int
	var username, password, email string
	var userID int
	var isLoggedIn bool

	for {
		if !isLoggedIn {
			// Пользователь сначала должен выбрать между логином и регистрацией
			fmt.Println("1. Login")
			fmt.Println("2. Register")
			fmt.Println("0. Exit")
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				// Логин
				fmt.Print("Enter username: ")
				fmt.Scan(&username)
				fmt.Print("Enter password: ")
				fmt.Scan(&password)
				if login(username, password) {
					isLoggedIn = true // Успешный логин
				} else {
					fmt.Println("Login failed. Please try again.")
				}
			case 2:
				// Регистрация
				fmt.Print("Enter username: ")
				fmt.Scan(&username)
				fmt.Print("Enter password: ")
				fmt.Scan(&password)
				fmt.Print("Enter email: ")
				fmt.Scan(&email)
				registerUser(username, password, email)
			case 0:
				// Выход из программы
				return
			default:
				fmt.Println("Invalid choice, please try again.")
			}
		} else {
			// После успешного входа доступен функционал для работы с пользователями
			fmt.Println("1. Get Users")
			fmt.Println("2. Update User")
			fmt.Println("3. Delete User")
			fmt.Println("4. Add User")
			fmt.Println("5. Logout") // Опция выхода из аккаунта
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				// Получить список пользователей
				getUsers()
			case 2:
				// Обновление пользователя
				fmt.Print("Enter user ID to update: ")
				fmt.Scan(&userID)
				fmt.Print("Enter new username: ")
				fmt.Scan(&username)
				fmt.Print("Enter new password: ")
				fmt.Scan(&password)
				fmt.Print("Enter new email: ")
				fmt.Scan(&email)
				updateUser(userID, username, password, email)
			case 3:
				// Удаление пользователя
				fmt.Print("Enter user ID to delete: ")
				fmt.Scan(&userID)
				deleteUser(userID)
			case 4:
				// Добавление пользователя
				fmt.Print("Enter username: ")
				fmt.Scan(&username)
				fmt.Print("Enter password: ")
				fmt.Scan(&password)
				fmt.Print("Enter email: ")
				fmt.Scan(&email)
				registerUser(username, password, email)
			case 5:
				// Выход из аккаунта
				fmt.Println("Logged out.")
				isLoggedIn = false
			default:
				fmt.Println("Invalid choice, please try again.")
			}
		}
	}
}

func deleteUser(id int) {
	user := User{ID: id}
	data, _ := json.Marshal(user)

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/users/delete", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	if sessionToken != "" {
		req.AddCookie(&http.Cookie{Name: "token", Value: sessionToken})
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("User with ID %d deleted successfully.\n", id)
	} else {
		handleResponse(resp) // Обработка ошибки
	}
}

func updateUser(id int, username, password, email string) {
	user := User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	data, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", "http://localhost:8080/users/update", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	if sessionToken != "" {
		req.AddCookie(&http.Cookie{Name: "token", Value: sessionToken})
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User updated:", string(body))
	} else {
		fmt.Printf("Updating user with ID: %d\n", user.ID)
		fmt.Println("Error:", string(body))
	}
}

func createUser(username, password, email string) {
	user := User{
		Username: username,
		Password: password,
		Email:    email,
	}
	data, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "http://localhost:8080/users/create", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	if sessionToken != "" {
		req.AddCookie(&http.Cookie{Name: "token", Value: sessionToken})
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User created:", string(body))
	} else {
		fmt.Println("Error:", string(body))
	}
}

func login(username, password string) bool {
	creds := Credentials{Username: username, Password: password}
	data, _ := json.Marshal(creds)

	resp, err := client.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return false // Возвращаем false при ошибке
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Сохраняем токен
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "token" {
				sessionToken = cookie.Value
			}
		}

		// Получаем ID текущего пользователя из ответа
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return false // Возвращаем false при ошибке
		}

		// Проверка, что ответ не пустой
		if len(body) == 0 {
			fmt.Println("Login response is empty.")
			return false // Возвращаем false, если тело ответа пустое
		}

		var loggedInUser User
		if err := json.Unmarshal(body, &loggedInUser); err == nil {
			currentUserID = loggedInUser.ID // сохраняем ID текущего пользователя
			fmt.Println("Logged in successfully.")
			return true // Возвращаем true при успешном входе
		} else {
			fmt.Println("Error parsing login response:", err)
			return false // Возвращаем false при ошибке парсинга
		}
	} else {
		fmt.Println("Failed to log in. Status code:", resp.StatusCode)
		return false // Возвращаем false при неуспешном ответе сервера
	}
}

func registerUser(username, password, email string) {
	user := User{
		Username: username,
		Password: password,
		Email:    email,
	}
	data, _ := json.Marshal(user)

	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("User successfully registered")
	} else if resp.StatusCode == http.StatusConflict {
		fmt.Println("Username already exists")
	} else {
		fmt.Println("Error:", resp.StatusCode)
	}
}

func getUsers() {
	req, _ := http.NewRequest("GET", "http://localhost:8080/users", nil)

	if sessionToken != "" {
		req.AddCookie(&http.Cookie{Name: "token", Value: sessionToken})
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		var users []User
		body, _ := ioutil.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &users); err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		if len(users) == 0 {
			fmt.Println("No users found.")
		} else {
			printUsers(users)
		}
	} else {
		handleResponse(resp)
	}
}

func printUsers(users []User) {
	fmt.Println("ID\tUsername\tpassword\tEmail")
	fmt.Println("-----------------------------------------------------------")
	for _, user := range users {
		fmt.Printf("%d\t%s\t\t%s\t\t%s\n", user.ID, user.Username, user.Password, user.Email)
	}
}

func handleResponse(resp *http.Response) {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Response: ", string(body))
	} else {
		fmt.Printf("Error %d: %s\n", resp.StatusCode, string(body))
	}
}

func testMultipleClients() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ { // 10 клиентов одновременно
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			fmt.Printf("Client %d trying to login...\n", clientID)
			login(fmt.Sprintf("user%d", clientID), "password123")
			getUsers()
		}(i)
	}
	wg.Wait()
}
