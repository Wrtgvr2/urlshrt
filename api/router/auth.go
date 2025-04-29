package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
	"github.com/wrtgvr/urlshrt/api/middleware"
)

func registerAuthRouter(h *handlers.Handler, r *gin.Engine) {
	group := r.Group("auth")
	{
		group.POST("/login", h.LoginHandler)
		group.POST("/register", h.RegisterHandler)
		group.POST("/refresh", middleware.AuthMiddleware(), h.RefreshTokenHandler)
	}
}
