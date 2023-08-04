package service

import "github.com/ZhanserikKalmukhambet/Trello/internal/entity"

func (s *TodoListService) CreateTodoList(userID int, list entity.TodoList) (int, error) {
	return s.repos.CreateTodoList(userID, list)
}

func (s *TodoListService) GetTodoLists(userID int) ([]entity.TodoList, error) {
	return s.repos.GetTodoLists(userID)
}

func (s *TodoListService) GetTodoListByID(userID, listID int) (entity.TodoList, error) {
	return s.repos.GetTodoListByID(userID, listID)
}

func (s *TodoListService) DeleteTodoList(userID, listID int) error {
	return s.repos.DeleteTodoList(userID, listID)
}

func (s *TodoListService) UpdateTodoList(userID, listID int, input entity.UpdateListInput) error {
	return s.repos.UpdateTodoList(userID, listID, input)
}
