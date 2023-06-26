package repository

import "github.com/RyushiAok/todo-go/domain/entities"

type IUserRepository interface {
	Create(name string) (*entities.User, error)
	Get(id string) (*entities.User, error)
	Delete(id string) error
}
