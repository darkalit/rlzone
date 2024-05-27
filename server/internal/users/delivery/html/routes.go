package html

import "github.com/gin-gonic/gin"

func MapItemRoutes(router *gin.RouterGroup, h *Handler) {
	itemsRoute := router.Group("/users")
	{
		itemsRoute.GET("/login", h.LoginGet)
		itemsRoute.POST("/login", h.LoginPost)
	}
}
