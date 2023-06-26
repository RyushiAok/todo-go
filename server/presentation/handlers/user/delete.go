package user_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type DeleteUserRequest struct {
	Id string `json:"id"`
}

func DeleteUser(repo *persistence.UserRepository, ctx *gin.Context) {
	var input DeleteUserRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := input.Id
	err := repo.Delete(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "deleted"},
	)
}
