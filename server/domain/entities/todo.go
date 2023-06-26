package entities

type Todo struct {
	Id          string `json:"todo_id" db:"todo_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
