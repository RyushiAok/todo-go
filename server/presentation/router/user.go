package router

import (
	"github.com/RyushiAok/todo-go/infrastructure/persistence"
	"github.com/RyushiAok/todo-go/presentation/handlers"
)

func (r Router) InitUserRouter() {
	repo := persistence.NewUserRepository()
	h := handlers.NewUserHandler(repo)
	g := r.Engine.Group("/user")
	g.POST("/", h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.DELETE("/:id", h.DeleteUser)
}
