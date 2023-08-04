package service

import "github.com/ZhanserikKalmukhambet/Trello/internal/entity"

func (s *TodoListService) CreateTodoList(userID int, list entity.TodoList) (int, error) {
	return s.repos.CreateTodoList(userID, list)
}

func (s *TodoListService) GetTodoLists(userID int) ([]entity.TodoList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TodoListService) GetTodoListByID(userID, listID int) (entity.TodoList, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TodoListService) DeleteTodoList(userID, listID int) error {
	//TODO implement me
	panic("implement me")
}

func (s *TodoListService) UpdateTodoList(userID, listID int, input entity.UpdateListInput) error {
	//TODO implement me
	panic("implement me")
}
