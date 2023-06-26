package repository

import "github.com/RyushiAok/todo-go/domain/entities"

type ITodoRepository interface {
	Create(title string, description string) (*entities.Todo, error)
	Get(id string) (*entities.Todo, error)
	Delete(id string) error
}
