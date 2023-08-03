package entity

type User struct {
	ID        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username" binding:"required"`
	Email     string     `json:"email" db:"email" binding:"required"`
	Password  string     `json:"password" db:"password" binding:"required"`
	TodoLists []TodoList `json:"todo_lists" db:"todo_lists"`
}

type RegisterRequest struct {
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
