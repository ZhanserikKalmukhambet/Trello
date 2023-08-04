package repository

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
)

func (r *TodoListPostgres) CreateTodoList(userID int, list entity.TodoList) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListPostgres) GetTodoLists(userID int) ([]entity.TodoList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListPostgres) GetTodoListByID(userID, listID int) (entity.TodoList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListPostgres) DeleteTodoList(userID, listID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *TodoListPostgres) UpdateTodoList(userID, listID int, input entity.UpdateListInput) error {
	//TODO implement me
	panic("implement me")
}
