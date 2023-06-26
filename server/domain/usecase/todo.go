package usecase

import (
	"fmt"

	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type ITodoUsecase interface {
	Create(repo *repository.ITodoRepository, title string, description string) (*entities.Todo, error)
	Get(repo *repository.ITodoRepository, id string) (*entities.Todo, error)
	Delete(repo *repository.ITodoRepository, id string) (*entities.Todo, error)
}

type TodoUsecase struct {
	repo repository.ITodoRepository
}

func NewTodoUsecase(repo repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (uc *TodoUsecase) Create(todo *entities.Todo) (*entities.Todo, error) {
	if todo.Title == "" {
		return nil, fmt.Errorf("title is empty")
	}
	if todo.Description == "" {
		return nil, fmt.Errorf("description is empty")
	}

	return uc.repo.Create(todo.Title, todo.Description)
}

func (uc *TodoUsecase) Get(id string) (*entities.Todo, error) {
	return uc.repo.Get(id)
}

func (uc *TodoUsecase) Delete(id string) error {
	err := uc.repo.Delete(id)
	return err
}
