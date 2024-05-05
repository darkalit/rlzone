package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/health"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	e.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	v1 := e.Group("/api/v1")

	health.MapHealthRoutes(v1)

	return nil
}
