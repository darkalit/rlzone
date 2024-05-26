package html

import "github.com/gin-gonic/gin"

func MapItemRoutes(router *gin.RouterGroup, h *Handler) {
	itemsRoute := router.Group("/items")
	{
		itemsRoute.GET("", h.Get)
	}
}
