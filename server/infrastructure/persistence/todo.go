package persistence

import (
	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type TodoRepository struct {
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

var _ repository.ITodoRepository = &TodoRepository{}

func (repo *TodoRepository) Create(title string, description string) (*entities.Todo, error) {
	todo := entities.Todo{
		Id:          "1",
		Title:       title,
		Description: description,
	}
	return &todo, nil
}

func (repo *TodoRepository) Delete(id string) error {
	return nil
}

func (repo *TodoRepository) Get(id string) (*entities.Todo, error) {
	todo := entities.Todo{
		Id:          id,
		Title:       "test",
		Description: "test",
	}
	return &todo, nil
}
