package todo_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

func GetTodo(uc usecase.ITodoUsecase, ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := uc.Get(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": todo},
	)
}
