package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type GetTodoRequest struct {
	Id string `json:"todo_id"`
}

func GetTodo(repo *persistence.TodoRepository, ctx *gin.Context) {
	var input GetTodoRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := repo.Get(input.Id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": todo},
	)
}
