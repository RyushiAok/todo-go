package usecase

import (
	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type IUserUsecase interface {
	Create(name string) (*entities.User, error)
	Get(id string) (*entities.User, error)
	Delete(id string) error
}

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(repo repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

var _ IUserUsecase = &UserUsecase{}

func (uu *UserUsecase) Create(name string) (*entities.User, error) {
	return uu.repo.Create(name)
}

func (uu *UserUsecase) Get(id string) (*entities.User, error) {
	return uu.repo.Get(id)
}

func (uu *UserUsecase) Delete(id string) error {
	return uu.repo.Delete(id)
}
