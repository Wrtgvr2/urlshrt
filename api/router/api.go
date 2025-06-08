package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
	"github.com/wrtgvr/urlshrt/api/middleware"
)

func registerApiRoutes(h *handlers.Handler, r *gin.Engine) {
	apiGroup := r.Group("api", middleware.AuthMiddleware())

	apiGroup.POST("/shorten", h.ShortenHandler)
	{
		usersGroup := apiGroup.Group("users")
		{
			usersGroup.GET("", h.GetUserHandler)
			usersGroup.PATCH("", h.PatchUserHandler)
			usersGroup.DELETE("", h.DeleteUserHandler)
		}
		urlsGroup := apiGroup.Group("urls")
		{
			urlsGroup.GET("", h.GetUserUrlsHandler)
			urlsGroup.GET("/:id", h.GetUrlHandler)
			urlsGroup.DELETE("/:id", h.DeleteUrlHandler)
		}
	}
}
