package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/pkg/db/mysql"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
	db     *gorm.DB
}

func NewServer() (*Server, error) {
	engine := gin.Default()

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	db, err := mysql.NewMySqlDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		engine,
		cfg,
		db,
	}, nil
}

func (s *Server) Run() error {
	err := s.MapHandlers(s.engine)
	if err != nil {
		return err
	}

	err = s.engine.Run(":" + s.config.AppPort)
	if err != nil {
		return err
	}
	return nil
}
