package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/health"
	"github.com/darkalit/rlzone/server/internal/items"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	e.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	v1 := e.Group("/api/v1")

	healthHandler := health.NewHandler()
	health.MapHealthRoutes(v1, healthHandler)

	itemsRepo := items.NewItemRepository(s.db)
	itemsUseCase := items.NewItemUseCase(itemsRepo)
	itemsHandler := items.NewHandler(s.config, itemsUseCase)
	items.MapItemRoutes(v1, itemsHandler)

	return nil
}
