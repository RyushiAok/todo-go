package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	e := gin.Default()
	return &Router{Engine: e}
}

func (r *Router) Run(port string) {
	err := r.Engine.Run(fmt.Sprintf(": %s", port))
	if err != nil {
		panic(err)
	}
}
