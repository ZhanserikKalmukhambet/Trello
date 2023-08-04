package service

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository"
)

type Authorization interface {
	Register(user entity.User) (int, error) // (newUserID, error )
	GetUserByID(id int) (entity.User, error)
	GenerateToken(username, password string) (string, error)
	VerifyToken(token string) (int, error)
}

type TodoList interface {
	CreateTodoList(userID int, list entity.TodoList) (int, error)
	GetTodoLists(userID int) ([]entity.TodoList, error)
	GetTodoListByID(userID, listID int) (entity.TodoList, error)
	DeleteTodoList(userID, listID int) error
	UpdateTodoList(userID, listID int, input entity.UpdateListInput) error
}

type TodoItem interface {
	CreateTodoItem(item entity.TodoItem) (int, error)
	GetTodoItems(listID int) ([]entity.TodoItem, error)
	DeleteTodoItem(listID, itemID int) error
	UpdateTodoItem(listID, itemID int, input entity.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repo.TodoItem),
	}
}
