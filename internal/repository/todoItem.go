package repository

import "github.com/ZhanserikKalmukhambet/Trello/internal/entity"

func (r *TodoItemPostgres) CreateTodoItem(listID int, item entity.TodoItem) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemPostgres) GetTodoItems(userID, listID int) ([]entity.TodoItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemPostgres) GetTodoItemByID(userID, itemID int) (entity.TodoItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemPostgres) DeleteTodoItem(userID, itemID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *TodoItemPostgres) UpdateTodoItem(userID, itemID int, input entity.UpdateItemInput) error {
	//TODO implement me
	panic("implement me")
}
