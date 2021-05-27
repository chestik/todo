package service

import (
	"crypto/sha1"
	"fmt"
	"rest_todo"
	"rest_todo/pkg/repository"
)

const salt = "qwoprueslkjdga.,vmxnzvc.b,mnaoihsl;a" // Как ключ для формирования hash

type AuthService struct {
	repo repository.Authorization
}

// NewAuthService Конструктор для структуры AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser Надстройка над структурой AuthService с функцией для создания пользователя
func (s *AuthService) CreateUser(user rest_todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// generatePasswordHash Функция генерации hash пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()           // Инициализация hash алгоритмом sha1
	hash.Write([]byte(password)) // Приведение к массиву байтов

	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) // Добавление "соли" для hash
}
