package main

import (
	"github.com/RyushiAok/todo-go/presentation/router"
)

func main() {
	r := router.NewRouter()
	r.Cors()
	r.InitHealth()
	r.InitUserRouter()
	r.InitTodoRouter()
	r.Run("8080")
}
