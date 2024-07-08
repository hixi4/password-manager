package password

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// представляє менеджер паролів, який зберігає паролі у файлі
type Manager struct {
	filePath  string            // Шлях до файлу для зберігання паролів
	passwords map[string]string // Карта для зберігання паролів
}

// створює новий екземпляр Manager і завантажує паролі з файлу
func NewManager(filePath string) (*Manager, error) {
	manager := &Manager{
		filePath:  filePath,
		passwords: make(map[string]string),
	}
	err := manager.loadPasswords()
	if err != nil {
		return nil, err
	}
	return manager, nil
}

// завантажує паролі з файлу
func (m *Manager) loadPasswords() error {
	if _, err := os.Stat(m.filePath); os.IsNotExist(err) {
		return nil // Якщо файл не існує, просто повертаємо nil
	}
	file, err := ioutil.ReadFile(m.filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &m.passwords)
	if err != nil {
		return err
	}
	return nil
}

// зберігає паролі у файл
func (m *Manager) savePasswords() error {
	data, err := json.Marshal(m.passwords)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(m.filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// виводить назви збережених паролів
func (m *Manager) ListPasswords() {
	fmt.Println("Saved passwords:")
	for name := range m.passwords {
		fmt.Println(name)
	}
}

// SavePassword зберігає новий пароль під заданою назвою
func (m *Manager) SavePassword(name, password string) error {
	if name == "" || password == "" {
		return errors.New("name and password cannot be empty")
	}
	m.passwords[name] = password
	return m.savePasswords()
}

// GetPassword повертає пароль за заданою назвою
func (m *Manager) GetPassword(name string) (string, error) {
	password, exists := m.passwords[name]
	if !exists {
		return "", errors.New("password not found")
	}
	return password, nil
}
