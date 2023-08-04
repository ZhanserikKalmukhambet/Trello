package service

import "github.com/ZhanserikKalmukhambet/Trello/internal/entity"

func (s *TodoItemService) CreateTodoItem(item entity.TodoItem) (int, error) {
	return s.repos.CreateTodoItem(item)
}

func (s *TodoItemService) GetTodoItems(listID int) ([]entity.TodoItem, error) {
	return s.repos.GetTodoItems(listID)
}

func (s *TodoItemService) DeleteTodoItem(listID, itemID int) error {
	return s.repos.DeleteTodoItem(listID, itemID)
}

func (s *TodoItemService) UpdateTodoItem(listID, itemID int, input entity.UpdateItemInput) error {
	return s.repos.UpdateTodoItem(listID, itemID, input)
}
