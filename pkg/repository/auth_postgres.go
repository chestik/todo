package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest_todo"
)

type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres Конструктор для инициализации структуры AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser Надстройка над структурой AuthPostgres с функцией для создания пользователя
func (r *AuthPostgres) CreateUser(user rest_todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
