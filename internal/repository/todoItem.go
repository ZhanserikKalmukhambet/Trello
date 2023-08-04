package repository

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository/pgrepo"
	"log"
)

func (r *TodoItemPostgres) CreateTodoItem(item entity.TodoItem) (int, error) {
	tableCreationQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    								id SERIAL PRIMARY KEY,
    								title VARCHAR(255),
    								description VARCHAR(255),
    								completed BOOLEAN DEFAULT FALSE,
    								todo_list_id INTEGER,
    								FOREIGN KEY (todo_list_id) REFERENCES %s(id))`,
		pgrepo.TodoItemsTable, pgrepo.TodoListsTable)
	_, err := r.db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, todo_list_id) VALUES ($1, $2, $3) RETURNING id", pgrepo.TodoItemsTable)

	row := r.db.QueryRow(query, item.Title, item.Description, item.TodoListID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TodoItemPostgres) GetTodoItems(listID int) ([]entity.TodoItem, error) {
	var items []entity.TodoItem

	query := fmt.Sprintf("SELECT * FROM %s WHERE todo_list_id=$1", pgrepo.TodoItemsTable)

	err := r.db.Select(&items, query, listID)

	return items, err
}

func (r *TodoItemPostgres) DeleteTodoItem(listID, itemID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND todo_list_id=$2", pgrepo.TodoItemsTable)

	_, err := r.db.Exec(query, itemID, listID)
	if err != nil{
		log.Fatal(err)
		return err
	}

	return nil
}

func (r *TodoItemPostgres) UpdateTodoItem(listID, itemID int, input entity.UpdateItemInput) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, completed=$3 WHERE id=$4 AND todo_list_id=$5", pgrepo.TodoItemsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, input.Completed, itemID, listID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
