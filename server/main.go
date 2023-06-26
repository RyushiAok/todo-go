package main

import (
	"github.com/RyushiAok/todo-go/presentation/router"
)

func main() {
	r := router.NewRouter()
	r.Cors()
	r.InitHealth()
	r.Run("8080")
}
