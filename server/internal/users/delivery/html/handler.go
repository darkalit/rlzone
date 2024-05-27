package html

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type Handler struct {
	cfg     *config.Config
	useCase users.UseCase
}

func NewHandler(cfg *config.Config, useCase users.UseCase) *Handler {
	return &Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (h *Handler) LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h *Handler) LoginPost(c *gin.Context) {
	var loginRequest users.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	foundUser, err := h.useCase.Login(c, &loginRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.SetCookie(
		"jwt",
		foundUser.RefreshToken,
		h.cfg.AuthCookieMaxAge,
		"/",
		c.Request.Host,
		h.cfg.AuthCookieSecure,
		h.cfg.AuthCookieHttpOnly,
	)

	c.Redirect(http.StatusFound, "/items")
}
