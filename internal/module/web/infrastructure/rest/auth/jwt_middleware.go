package auth

import (
	"net/http"
	"place4live/internal/module/web/app/port"
	"strings"

	"github.com/gin-gonic/gin"
)

const headerAuthorization = "Authorization"

func JwtAuthMiddleware(inPort port.JwtTokenQueryInPort) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := inPort.Get(getToken(c))
		if err != nil {
			c.String(http.StatusBadRequest, "Missed JWT token")
			c.Abort()
			return
		}

		c.Set("user_id", token.UserId)
		c.Next()
	}
}

func getToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get(headerAuthorization)
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
