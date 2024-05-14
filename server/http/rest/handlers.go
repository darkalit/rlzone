package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/health"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/internal/middleware"
	"github.com/darkalit/rlzone/server/internal/routes"
	"github.com/darkalit/rlzone/server/internal/users"
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

	healthHandler := health.NewHandler()

	usersRepo := users.NewUserRepository(s.config, s.db)
	usersUseCase := users.NewUserUseCase(usersRepo, s.config)
	usersHandler := users.NewHandler(s.config, usersUseCase)

	itemsRepo := items.NewItemRepository(s.config, s.db)
	itemsUseCase := items.NewItemUseCase(itemsRepo)
	itemsHandler := items.NewHandler(s.config, itemsUseCase)

	mw := middleware.NewMiddlewareManager(s.config, usersUseCase)
	routes.MapHealthRoutes(v1, healthHandler)
	routes.MapUserRoutes(v1, usersHandler, mw)
	routes.MapItemRoutes(v1, itemsHandler, mw)

	return nil
}
