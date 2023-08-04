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
	GetTodoListByID(listID int) (entity.TodoList, error)
	DeleteTodoList(userID, listID int) error
	UpdateTodoList(listID int, input entity.UpdateListInput) error
}

type TodoItem interface {
	CreateTodoItem(listID int, item entity.TodoItem) (int, error)
	GetTodoItems(userID, listID int) ([]entity.TodoItem, error)
	GetTodoItemByID(userID, itemID int) (entity.TodoItem, error)
	DeleteTodoItem(userID, itemID int) error
	UpdateTodoItem(userID, itemID int, input entity.UpdateItemInput) error
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
