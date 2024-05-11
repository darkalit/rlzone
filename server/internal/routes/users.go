package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/middleware"
	"github.com/darkalit/rlzone/server/internal/users"
)

func MapUserRoutes(router *gin.RouterGroup, h *users.Handler, mw *middleware.MiddlewareManager) {
	usersRoute := router.Group("/users")
	{
		usersRoute.POST("/register", h.Register)
		usersRoute.POST("/login", h.Login)
		usersRoute.GET("/refresh", h.Refresh)
		usersRoute.GET("/logout", h.Logout)

		usersRoute.GET("/block/:id", mw.AuthJWTMiddleware, mw.PermitAdmin, h.BlockUser)
	}
}
