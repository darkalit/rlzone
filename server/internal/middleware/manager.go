package middleware

import (
	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/users"
)

type MiddlewareManager struct {
	cfg     *config.Config
	usersUC users.UseCase
}

func NewMiddlewareManager(cfg *config.Config, usersUC users.UseCase) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:     cfg,
		usersUC: usersUC,
	}
}
