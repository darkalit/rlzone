package middleware

import (
	"github.com/darkalit/rlzone/server/config"
)

type MiddlewareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *MiddlewareManager {
	return &MiddlewareManager{
		cfg: cfg,
	}
}
