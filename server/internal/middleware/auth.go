package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/auth"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

func (mw *MiddlewareManager) SetPayload(c *gin.Context) {
	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		c.Next()
		return
	}

	payload, err := auth.VerifyToken(jwtCookie, mw.cfg, auth.RefreshTokenType)
	if err != nil {
		c.Next()
		return
	}

	c.Set("payload", *payload)
	c.Next()
}

func (mw *MiddlewareManager) AuthJWTMiddleware(c *gin.Context) {
	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(
			httpErrors.ErrorResponse(
				httpErrors.NewRestError(http.StatusUnauthorized, "JWT is undefined", err),
			),
		)
		return
	}

	payload, err := auth.VerifyToken(jwtCookie, mw.cfg, auth.RefreshTokenType)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	foundUser, err := mw.usersUC.GetById(c, payload.UserID)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	if foundUser.IsBlocked {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	c.Set("payload", *payload)
	c.Next()
}

func (mw *MiddlewareManager) PermitAdmin(c *gin.Context) {
	payloadAny, ok := c.Get("payload")
	if !ok {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}
	payload := payloadAny.(auth.JWTPayload)

	if payload.Role != string(users.RoleAdmin) {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	c.Next()
}
