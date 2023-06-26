package handlers

import (
	"github.com/RyushiAok/todo-go/domain/usecase"
	user_handler "github.com/RyushiAok/todo-go/presentation/handlers/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.IUserUsecase
}

func NewUserHandler(uc usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	user_handler.CreateUser(u.uc, ctx)

}

func (u *UserHandler) GetUser(ctx *gin.Context) {
	user_handler.GetUser(u.uc, ctx)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	user_handler.DeleteUser(u.uc, ctx)
}
