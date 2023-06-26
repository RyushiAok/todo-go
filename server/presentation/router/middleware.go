package router

import (
	"time"

	"github.com/gin-contrib/cors"
)

func (r *Router) Cors() {
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"OPTIONS",
			"PUT",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}
