package user_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type GetUserRequest struct {
	Id string `json:"id"`
}

func GetUser(repo *persistence.UserRepository, ctx *gin.Context) {
	var input GetUserRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := input.Id
	user, err := repo.Get(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": user},
	)
}
