package user_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

type DeleteUserRequest struct {
	Id string `json:"id"`
}

func DeleteUser(uc usecase.IUserUsecase, ctx *gin.Context) {
	var input DeleteUserRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := input.Id
	err := uc.Delete(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "deleted"},
	)
}
