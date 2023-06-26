package usecase

import (
	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type ITodoUsecase interface {
	Create(title string, description string) (*entities.Todo, error)
	Get(id string) (*entities.Todo, error)
	Delete(id string) error
}

type TodoUsecase struct {
	repo repository.ITodoRepository
}

func NewTodoUsecase(repo repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

var _ ITodoUsecase = &TodoUsecase{}

func (uc *TodoUsecase) Create(title, description string) (*entities.Todo, error) {
	return uc.repo.Create(title, description)
}

func (uc *TodoUsecase) Get(id string) (*entities.Todo, error) {
	return uc.repo.Get(id)
}

func (uc *TodoUsecase) Delete(id string) error {
	err := uc.repo.Delete(id)
	return err
}
