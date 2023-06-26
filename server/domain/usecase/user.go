package usecase

import (
	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type IUserUsecase interface {
	CreateUser(repo *repository.IUserRepository, name string) (*entities.User, error)
	GetUser(repo *repository.IUserRepository, id string) (*entities.User, error)
	DeleteUser(repo *repository.IUserRepository, id string) (*entities.User, error)
}

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(repo repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uu *UserUsecase) Create(name string) (*entities.User, error) {
	return uu.repo.Create(name)
}

func (uu *UserUsecase) Get(id string) (*entities.User, error) {
	return uu.repo.Get(id)
}

func (uu *UserUsecase) Delete(id string) error {
	return uu.repo.Delete(id)
}
