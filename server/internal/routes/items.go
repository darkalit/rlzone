package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/internal/middleware"
)

func MapItemRoutes(router *gin.RouterGroup, h *items.Handler, mw *middleware.MiddlewareManager) {
	itemsRoute := router.Group("/items")
	{
		itemsRoute.GET("", h.Get)
		itemsRoute.GET(":id", h.GetById)

		itemsRoute.POST("/stocks", mw.AuthJWTMiddleware, mw.PermitAdmin, h.CreateStock)
	}
}
