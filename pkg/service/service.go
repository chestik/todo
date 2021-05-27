package service

import (
	"rest_todo"
	"rest_todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest_todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      nil,
		TodoItem:      nil,
	}
}
