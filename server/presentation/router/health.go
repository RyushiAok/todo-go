package router

import (
	handler "github.com/RyushiAok/todo-go/presentation/handlers"
)

func (r Router) InitHealth() {
	r.Engine.GET("/health", handler.HealthHandler)
}
