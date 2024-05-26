package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	restHealth "github.com/darkalit/rlzone/server/internal/health/delivery/rest"
	"github.com/darkalit/rlzone/server/internal/items"
	htmlItems "github.com/darkalit/rlzone/server/internal/items/delivery/html"
	restItems "github.com/darkalit/rlzone/server/internal/items/delivery/rest"
	"github.com/darkalit/rlzone/server/internal/middleware"
	"github.com/darkalit/rlzone/server/internal/users"
	restUsers "github.com/darkalit/rlzone/server/internal/users/delivery/rest"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	e.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}),
	)

	v1 := e.Group("/api/v1")
	h := e.Group("/")

	healthRestHandler := restHealth.NewHandler()

	usersRepo := users.NewUserRepository(s.config, s.db)
	usersUseCase := users.NewUserUseCase(usersRepo, s.config)
	usersRestHandler := restUsers.NewHandler(s.config, usersUseCase)

	itemsRepo := items.NewItemRepository(s.config, s.db)
	itemsUseCase := items.NewItemUseCase(itemsRepo)
	itemsRestHandler := restItems.NewHandler(s.config, itemsUseCase)
	itemsHtmlHandler := htmlItems.NewHandler(s.config, itemsUseCase)

	mw := middleware.NewMiddlewareManager(s.config, usersUseCase)
	restHealth.MapHealthRoutes(v1, healthRestHandler)
	restUsers.MapUserRoutes(v1, usersRestHandler, mw)
	restItems.MapItemRoutes(v1, itemsRestHandler, mw)
	htmlItems.MapItemRoutes(h, itemsHtmlHandler)

	return nil
}
