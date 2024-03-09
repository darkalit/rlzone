package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	e.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	v1 := e.Group("/api/v1")

	health := v1.Group("/health")

	health.GET("", func(
		ctx *gin.Context,
	) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	return nil
}
