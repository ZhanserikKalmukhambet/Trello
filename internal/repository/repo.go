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
	CreateTodoItem(item entity.TodoItem) (int, error)
	GetTodoItems(listID int) ([]entity.TodoItem, error)
	DeleteTodoItem(listID, itemID int) error
	UpdateTodoItem(listID, itemID int, input entity.UpdateItemInput) error
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
