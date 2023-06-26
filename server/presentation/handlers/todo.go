package handlers

import (
	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	todo_handler "github.com/RyushiAok/todo-go/presentation/handlers/todo"
	"github.com/gin-gonic/gin"
)

type TodoHandlers struct {
	repo *persistence.TodoRepository
}

func NewTodoHandlers(repo *persistence.TodoRepository) *TodoHandlers {
	return &TodoHandlers{
		repo: repo,
	}
}

func (u *TodoHandlers) CreateTodo(ctx *gin.Context) {
	todo_handler.CreateTodo(u.repo, ctx)

}

func (u *TodoHandlers) GetTodo(ctx *gin.Context) {
	todo_handler.GetTodo(u.repo, ctx)
}

func (u *TodoHandlers) DeleteTodo(ctx *gin.Context) {
	todo_handler.DeleteTodo(u.repo, ctx)
}
