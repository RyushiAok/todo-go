package entities

type User struct {
	Id   string `json:"user_id" db:"user_id"`
	Name string `json:"name" db:"name"`
}
