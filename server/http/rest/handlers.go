package rest

import (
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
	)

	v1 := e.Group("/api/v1")

	healthHandler := health.NewHandler()

	usersRepo := users.NewUserRepository(s.db)
	usersUseCase := users.NewUserUseCase(usersRepo, s.config)
	usersHandler := users.NewHandler(s.config, usersUseCase)

	itemsRepo := items.NewItemRepository(s.db)
	itemsUseCase := items.NewItemUseCase(itemsRepo)
	itemsHandler := items.NewHandler(s.config, itemsUseCase)

	mw := middleware.NewMiddlewareManager(s.config, usersUseCase)
	routes.MapHealthRoutes(v1, healthHandler)
	routes.MapUserRoutes(v1, usersHandler, mw)
	routes.MapItemRoutes(v1, itemsHandler, mw)

	return nil
}
