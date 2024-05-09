package health

import (
	"github.com/gin-gonic/gin"
)

func MapHealthRoutes(router *gin.RouterGroup, h *Handler) {
	health := router.Group("/health")
	{
		health.GET("", h.Get)
	}
}
