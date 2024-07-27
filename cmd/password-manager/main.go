package main

import (
	"fmt"
	"os"
	"password-manager/internal/password"
)

func main() {
	// Створюємо новий менеджер паролів, який буде зберігати дані у файлі passwords.json
	manager, err := password.NewManager("passwords.json")
	if err != nil {
		fmt.Println("Error initializing password manager:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./binary_name <command> [<args>]")
		fmt.Println("Commands:")
		fmt.Println("  list                List saved passwords")
		fmt.Println("  put <name> <pass>   Save a new password")
		fmt.Println("  get <name>          Retrieve a saved password")
		return
	}

	switch os.Args[1] {
	case "list":
		manager.ListPasswords()
	case "put":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./binary_name put <name> <pass>")
			return
		}
		name := os.Args[2]
		password := os.Args[3]
		err := manager.SavePassword(name, password)
		if err != nil {
			fmt.Println("Error saving password:", err)
		}
	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: ./binary_name get <name>")
			return
		}
		name := os.Args[2]
		password, err := manager.GetPassword(name)
		if err != nil {
			fmt.Println("Error retrieving password:", err)
		} else {
			fmt.Println("Password:", password)
		}
	default:
		fmt.Println("Invalid command. Use one of the following:")
		fmt.Println("  list                List saved passwords")
		fmt.Println("  put <name> <pass>   Save a new password")
		fmt.Println("  get <name>          Retrieve a saved password")
	}
}
