package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type Handler struct {
	cfg     *config.Config
	useCase UseCase
}

func NewHandler(cfg *config.Config, useCase UseCase) *Handler {
	return &Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var registerRequest RegisterRequest

	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	createdUser, err := h.useCase.Register(c, &registerRequest)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.SetCookie(
		"jwt",
		createdUser.RefreshToken,
		h.cfg.AuthCookieMaxAge,
		"/",
		c.Request.Host,
		h.cfg.AuthCookieSecure,
		h.cfg.AuthCookieHttpOnly,
	)

	c.JSON(http.StatusOK, createdUser)
}

func (h *Handler) Login(c *gin.Context) {
	var loginRequest LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
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

	c.JSON(http.StatusOK, foundUser)
}

func (h *Handler) Refresh(c *gin.Context) {
	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(
			httpErrors.ErrorResponse(
				httpErrors.NewRestError(http.StatusUnauthorized, "JWT is undefined", err),
			),
		)
		return
	}

	foundUser, err := h.useCase.RefreshToken(c, jwtCookie)
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

	c.JSON(http.StatusOK, foundUser)
}

func (h *Handler) Logout(c *gin.Context) {
	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err = h.useCase.Logout(c, jwtCookie)
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.SetCookie(
		"jwt",
		"",
		-1,
		"/",
		c.Request.Host,
		h.cfg.AuthCookieSecure,
		h.cfg.AuthCookieHttpOnly,
	)

	c.Status(http.StatusNoContent)
}

func (h *Handler) BlockUser(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	id := uint(id64)

	err = h.useCase.BlockUser(c, id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
