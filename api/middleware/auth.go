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
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		err := jwt.ValidateToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token", err, true))
			c.Abort()
			return
		}
		userId, err := jwt.GetUserIdFromToken(tokenStr)
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid token payload", err, true))
			c.Abort()
			return
		}
<<<<<<< HEAD
=======
		if err != nil {
			ginadap.HandleError(c, errsuit.NewUnauthorized("invalid userID in token", err, true))
			c.Abort()
			return
		}
>>>>>>> dfbaed329160dd3ea8b7f0e3fc79de58d3c80e63
		c.Set("UserID", userId)

		c.Next()
	}
}
