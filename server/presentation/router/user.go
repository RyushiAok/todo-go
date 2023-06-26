package router

import (
	"github.com/RyushiAok/todo-go/domain/usecase"
	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/RyushiAok/todo-go/presentation/handlers"
)

func (r Router) InitUserRouter() {
	repo := persistence.NewUserRepository()
	uc := usecase.NewUserUsecase(repo)
	h := handlers.NewUserHandler(uc)
	g := r.Engine.Group("/user")
	g.POST("/", h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.DELETE("/:id", h.DeleteUser)
}
