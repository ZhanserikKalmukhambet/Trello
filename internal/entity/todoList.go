package entity

type TodoList struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	TodoItems   []TodoItem `json:"todo_items" db:"todo_items"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
