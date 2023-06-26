package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

func DeleteTodo(uc usecase.ITodoUsecase, ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.Delete(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": "deleted"},
	)
}
