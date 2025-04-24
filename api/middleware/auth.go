package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/pkg/jwt"
	"github.com/wrtgvr2/errsuit"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ginadap.HandleError(c, errsuit.NewUnauthorized("missing authorization header", nil, false))
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		err := jwt.ValidateAccessToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			c.Abort()
			return
		}

		c.Next()
	}
}
