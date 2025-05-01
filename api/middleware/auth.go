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
			ginadap.HandleError(c, errsuit.NewUnauthorized("missing authorization header", nil, true))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		err := jwt.ParseToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			return
		}

		tokenType, err := jwt.GetTokenType(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			return
		}
		if tokenType != jwt.TypeAccess {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			return
		}

		userId, err := jwt.GetUserIdFromToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token payload", err, true))
			return
		}
		c.Set("UserID", userId)

		c.Next()
	}
}
