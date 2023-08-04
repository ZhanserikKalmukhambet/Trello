package entity

type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Completed   bool   `json:"completed" db:"completed"`
	TodoListID  int    `json:"todo_list_id" db:"todo_list_id"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}
