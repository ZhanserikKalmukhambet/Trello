package service

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"gorm.io/gorm"
)

type Authorization interface {
	Register(user entity.User) (int, error)
	Login(username, password string) (entity.User, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item entity.TodoItem) (int, error)
	GetAll(userId, listId int) ([]entity.TodoItem, error)
	GetById(userId, itemId int) (entity.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input entity.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(db *gorm.DB) *Service {
	return &Service{}
}
