package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/auth"
)

func (mw *MiddlewareManager) AuthJWTMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	t := strings.Split(authHeader, " ")
	if len(t) != 2 {
		c.Status(http.StatusForbidden)
		c.Abort()
		return
	}

	token := t[1]
	payload, err := auth.VerifyToken(token, mw.cfg, auth.AccessTokenType)
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
