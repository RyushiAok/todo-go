package user_handler

import (
	"net/http"

	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/gin-gonic/gin"
)

func GetUser(uc usecase.IUserUsecase, ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.Get(id)
	if err != nil {
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": user},
	)
}
