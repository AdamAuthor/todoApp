package repository

import (
	"todoApp/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listItemsTable  = "list_items"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string) (models.User, error)
}

type TodoList interface {
	Create(userID int, list models.List) (int, error)
	GetAll(userID int) ([]models.List, error)
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
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
	}
}
