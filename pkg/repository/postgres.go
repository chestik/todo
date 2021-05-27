package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListTable   = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	itemsListsTable = "items_lists"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB Конструктор для установления соедениния с бд согласно данным из cfg
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Установка соединения с БД
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	// Проверка "пинга" от бд
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
