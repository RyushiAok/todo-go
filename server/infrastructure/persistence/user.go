package persistence

import (
	"github.com/RyushiAok/todo-go/domain/entities"
	"github.com/RyushiAok/todo-go/domain/repository"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var _ repository.IUserRepository = &UserRepository{}

func (repo *UserRepository) Create(name string) (*entities.User, error) {
	user := entities.User{
		Id:   "1",
		Name: "test",
	}
	return &user, nil
}

func (repo *UserRepository) Delete(id string) error {
	return nil
}

func (repo *UserRepository) Get(id string) (*entities.User, error) {
	user := entities.User{
		Id:   "1",
		Name: "test",
	}
	return &user, nil
}
