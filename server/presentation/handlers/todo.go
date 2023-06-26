package handlers

import (
	"github.com/RyushiAok/todo-go/domain/usecase"
	todo_handler "github.com/RyushiAok/todo-go/presentation/handlers/todo"
	"github.com/gin-gonic/gin"
)

type TodoHandlers struct {
	uc usecase.ITodoUsecase
}

func NewTodoHandlers(uc usecase.ITodoUsecase) *TodoHandlers {
	return &TodoHandlers{
		uc: uc,
	}
}

func (u *TodoHandlers) CreateTodo(ctx *gin.Context) {
	todo_handler.CreateTodo(u.uc, ctx)

}

func (u *TodoHandlers) GetTodo(ctx *gin.Context) {
	todo_handler.GetTodo(u.uc, ctx)
}

func (u *TodoHandlers) DeleteTodo(ctx *gin.Context) {
	todo_handler.DeleteTodo(u.uc, ctx)
}
