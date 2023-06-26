package user_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

func CreateUser(uc usecase.IUserUsecase, ctx *gin.Context) {
	var input CreateUserRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := input.Name
	user, err := uc.Create(name)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": user},
	)
}
