package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

type DeleteTodoRequest struct {
	Id string `json:"todo_id"`
}

func DeleteTodo(repo *persistence.TodoRepository, ctx *gin.Context) {
	var input DeleteTodoRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.Delete(input.Id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "deleted"},
	)
}
