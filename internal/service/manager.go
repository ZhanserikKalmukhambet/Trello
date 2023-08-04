package service

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

/////////////////

type TodoListService struct {
	repos repository.TodoList
}

func NewTodoListService(repos repository.TodoList) *TodoListService {
	return &TodoListService{repos: repos}
}

//////////////////

type TodoItemService struct {
	repos repository.TodoItem
}

func NewTodoItemService(repos repository.TodoItem) *TodoItemService {
	return &TodoItemService{repos: repos}
}
