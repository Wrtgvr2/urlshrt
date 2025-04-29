package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
	"github.com/wrtgvr/urlshrt/api/middleware"
)

func registerApiRoutes(h *handlers.Handler, r *gin.Engine) {
	apiGroup := r.Group("api", middleware.AuthMiddleware())
	{
		rGroup := apiGroup.Group("r")
		{
			rGroup.GET("/:shrturl", h.RedirectHandler)
			//rGroup.POST("/shorten", h.ShortenHandler)
		}
		usersGroup := apiGroup.Group("users")
		{
			usersGroup.GET("/:id", h.GetUserHandler)
			// GET all
			// PATCH {id}
			// DELETE {id}
		}
		// urlsGroup := apiGroup.Group("urls")
		// {
		//   POST create
		//   POST regenerate short url with same orig url
		//   GET all
		//   GET {id}
		// }
	}
}
