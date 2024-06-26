package service

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list models.List) (int, error)
	GetAll(userID int) ([]models.List, error)
	GetByID(userID, listID int) (models.List, error)
	Update(userID, listID int, input models.UpdateListInput) error
	Delete(userID, listID int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
