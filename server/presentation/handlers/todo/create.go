package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTodo(uc usecase.ITodoUsecase, ctx *gin.Context) {
	var input CreateTodoRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := uc.Create(input.Title, input.Description)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": todo},
	)
}
