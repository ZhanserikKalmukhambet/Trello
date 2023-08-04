package repository

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository/pgrepo"
	"log"
)

func (r *TodoListPostgres) CreateTodoList(userID int, list entity.TodoList) (int, error) {
	tableCreationQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    								id SERIAL PRIMARY KEY,
    								title VARCHAR(255),
    								description VARCHAR(255),
    								user_id INTEGER,
    								FOREIGN KEY (user_id) REFERENCES %s(id))`,
		pgrepo.TodoListsTable, pgrepo.UsersTable)
	_, err := r.db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, user_id) VALUES ($1, $2, $3) RETURNING id", pgrepo.TodoListsTable)

	row := r.db.QueryRow(query, list.Title, list.Description, userID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TodoListPostgres) GetTodoLists(userID int) ([]entity.TodoList, error) {
	var lists []entity.TodoList

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", pgrepo.TodoListsTable)
	err := r.db.Select(&lists, query, userID)

	return lists, err
}

func (r *TodoListPostgres) GetTodoListByID(userID, listID int) (entity.TodoList, error) {
	var lists []entity.TodoList
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND user_id=$2", pgrepo.TodoListsTable)
	err := r.db.Select(&lists, query, listID, userID)

	return lists[0], err
}

func (r *TodoListPostgres) DeleteTodoList(userID, listID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", pgrepo.TodoListsTable)
	_, err := r.db.Exec(query, listID, userID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (r *TodoListPostgres) UpdateTodoList(userID, listID int, input entity.UpdateListInput) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3 AND user_id=$4", pgrepo.TodoListsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, listID, userID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
