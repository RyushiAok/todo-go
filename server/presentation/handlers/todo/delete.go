package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

type DeleteTodoRequest struct {
	Id string `json:"todo_id"`
}

func DeleteTodo(uc usecase.ITodoUsecase, ctx *gin.Context) {
	var input DeleteTodoRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.Delete(input.Id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "deleted"},
	)
}
