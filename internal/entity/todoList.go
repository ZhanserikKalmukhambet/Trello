package entity

type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	UserID      int    `json:"user_id" db:"user_id" binding:"required"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
