package repository

import (
	"github.com/jmoiron/sqlx"
	"rest_todo"
)

type Authorization interface {
	CreateUser(user rest_todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db)}
}
