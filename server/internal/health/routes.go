package health

import (
	"github.com/gin-gonic/gin"
)

func MapHealthRoutes(router *gin.RouterGroup) {
	h := NewHandler()

	health := router.Group("/health")
	{
		health.GET("", h.Get)
	}
}
