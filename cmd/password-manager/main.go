package main

import (
	"bufio"
	"fmt"
	"os"
	"password-manager/internal/password"
	"strings"
)

func main() {
	// Створюємо новий менеджер паролів, який буде зберігати дані у файлі passwords.json
	manager, err := password.NewManager("passwords.json")
	if err != nil {
		fmt.Println("Error initializing password manager:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin) // Створюємо сканер для читання введення з консолі

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. List saved passwords")
		fmt.Println("2. Save a new password")
		fmt.Println("3. Retrieve a saved password")
		fmt.Println("4. Exit")

		if !scanner.Scan() {
			fmt.Println("Error reading input.")
			continue
		}
		choice := strings.TrimSpace(scanner.Text()) // Читаємо вибір користувача

		switch choice {
		case "1":
			manager.ListPasswords() // Виводимо назви збережених паролів
		case "2":
			fmt.Print("Enter name: ")
			if !scanner.Scan() {
				fmt.Println("Error reading input.")
				continue
			}
			name := strings.TrimSpace(scanner.Text()) // Читаємо назву паролю

			fmt.Print("Enter password: ")
			if !scanner.Scan() {
				fmt.Println("Error reading input.")
				continue
			}
			password := strings.TrimSpace(scanner.Text()) // Читаємо сам пароль

			// Зберігаємо пароль під заданою назвою
			err := manager.SavePassword(name, password)
			if err != nil {
				fmt.Println("Error saving password:", err)
			}
		case "3":
			fmt.Print("Enter name: ")
			if !scanner.Scan() {
				fmt.Println("Error reading input.")
				continue
			}
			name := strings.TrimSpace(scanner.Text()) // Читаємо назву паролю для отримання

			// Отримуємо пароль за заданою назвою
			password, err := manager.GetPassword(name)
			if err != nil {
				fmt.Println("Error retrieving password:", err)
			} else {
				fmt.Println("Password:", password)
			}
		case "4":
			return // Вихід з програми
		default:
			fmt.Println("Invalid choice, please try again.") // Невірний вибір
		}
	}
}
