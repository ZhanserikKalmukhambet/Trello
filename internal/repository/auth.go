package repository

import (
	"fmt"
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository/pgrepo"
	"log"
)

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	tableCreationQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    								id SERIAL PRIMARY KEY,
    								username VARCHAR(255),
    								email VARCHAR(255),
    								password VARCHAR(255))`, pgrepo.UsersTable)
	_, err := r.db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) VALUES ($1, $2, $3) RETURNING id", pgrepo.UsersTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", pgrepo.UsersTable)
	err := r.db.Get(&user, query, username)

	return user, err
}

func (r *AuthPostgres) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", pgrepo.UsersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}
