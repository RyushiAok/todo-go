package router

import (
	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/RyushiAok/todo-go/presentation/handlers"
)

func (r Router) InitTodoRouter() {
	repo := persistence.NewTodoRepository()
	uc := usecase.NewTodoUsecase(repo)
	h := handlers.NewTodoHandlers(uc)
	g := r.Engine.Group("/todo")
	g.POST("/", h.CreateTodo)
	g.GET("/:id", h.GetTodo)
	g.DELETE("/:id", h.DeleteTodo)
}
