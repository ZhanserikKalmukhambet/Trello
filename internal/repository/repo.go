package repository

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username string) (entity.User, error)
	GetUserByID(userID int) (entity.User, error)
}

type TodoList interface {
	CreateTodoList(userID int, list entity.TodoList) (int, error)
	GetTodoLists(userID int) ([]entity.TodoList, error)
	GetTodoListByID(userID, listID int) (entity.TodoList, error)
	DeleteTodoList(userID, listID int) error
	UpdateTodoList(userID, listID int, input entity.UpdateListInput) error
}

type TodoItem interface {
	CreateTodoItem(listID int, item entity.TodoItem) (int, error)
	GetTodoItems(userID, listID int) ([]entity.TodoItem, error)
	GetTodoItemByID(userID, itemID int) (entity.TodoItem, error)
	DeleteTodoItem(userID, itemID int) error
	UpdateTodoItem(userID, itemID int, input entity.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
