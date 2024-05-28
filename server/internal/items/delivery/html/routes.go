package html

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/middleware"
)

func MapItemRoutes(router *gin.RouterGroup, h *Handler, mw *middleware.MiddlewareManager) {
	itemsRoute := router.Group("/items")
	{
		itemsRoute.GET("", h.Get)
		itemsRoute.GET("/create", h.CreateStockGet)
		itemsRoute.POST("/create", h.CreateStockPost)
		itemsRoute.GET("/list", h.GetList)
		itemsRoute.POST("/buy", mw.AuthJWTMiddleware, h.BuyItem)
	}
}
