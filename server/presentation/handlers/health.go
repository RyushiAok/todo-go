package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Health string `json:"health"`
}

func HealthHandler(ctx *gin.Context) {
	h := HealthResponse{Health: "ok"}
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": h},
	)
}
