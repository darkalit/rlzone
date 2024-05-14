package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/health"
)

func MapHealthRoutes(router *gin.RouterGroup, h *health.Handler) {
	health := router.Group("/health")
	{
		health.GET("", h.Get)
	}
}
