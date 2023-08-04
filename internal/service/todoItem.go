package service

import "github.com/ZhanserikKalmukhambet/Trello/internal/entity"

func (s *TodoItemService) CreateTodoItem(listID int, item entity.TodoItem) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TodoItemService) GetTodoItems(userID, listID int) ([]entity.TodoItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TodoItemService) GetTodoItemByID(userID, itemID int) (entity.TodoItem, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TodoItemService) DeleteTodoItem(userID, itemID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *TodoItemService) UpdateTodoItem(userID, itemID int, input entity.UpdateItemInput) error {
	//TODO implement me
	panic("implement me")
}
