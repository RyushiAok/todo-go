package handlers

import (
	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	user_handler "github.com/RyushiAok/todo-go/presentation/handlers/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo *persistence.UserRepository
}

func NewUserHandler(repo *persistence.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	user_handler.CreateUser(u.repo, ctx)

}

func (u *UserHandler) GetUser(ctx *gin.Context) {
	user_handler.GetUser(u.repo, ctx)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	user_handler.DeleteUser(u.repo, ctx)
}
