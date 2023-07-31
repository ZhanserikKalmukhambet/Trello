package entity

type User struct {
	ID        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"password"`
	TodoLists []TodoList `json:"todo_lists" db:"todo_lists"`
}
