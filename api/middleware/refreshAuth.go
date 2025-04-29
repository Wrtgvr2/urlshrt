package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/pkg/jwt"
	"github.com/wrtgvr2/errsuit"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func RefreshAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ginadap.HandleError(c, errsuit.NewUnauthorized("missing authorization header", nil, true))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		err := jwt.ValidateToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			return
		}

		jti, err := jwt.GetJtiFromRefreshToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token payload", err, true))
			return
		}
		c.Set("JTI", jti)

		userId, err := jwt.GetUserIdFromToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token payload", err, true))
			return
		}
		c.Set("UserID", userId)

		c.Next()
	}
}
