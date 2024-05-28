package html

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/middleware"
)

func MapItemRoutes(router *gin.RouterGroup, h *Handler, mw *middleware.MiddlewareManager) {
	itemsRoute := router.Group("/items")
	{
		itemsRoute.GET("", h.Get)
		itemsRoute.GET("/create", mw.AuthJWTMiddleware, mw.PermitAdmin, h.CreateStockGet)
		itemsRoute.POST("/create", mw.AuthJWTMiddleware, mw.PermitAdmin, h.CreateStockPost)
		itemsRoute.GET("/list", h.GetList)
		itemsRoute.POST("/buy", mw.AuthJWTMiddleware, h.BuyItem)
		itemsRoute.GET("/inventory", mw.AuthJWTMiddleware, h.Inventory)
		itemsRoute.POST("/sell", mw.AuthJWTMiddleware, h.SellItem)
	}
}
