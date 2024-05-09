package items

import "github.com/gin-gonic/gin"

func MapItemRoutes(router *gin.RouterGroup, h *Handler) {
	items := router.Group("/items")
	{
		items.GET("", h.Get)
		items.GET(":id", h.GetById)
	}
}
