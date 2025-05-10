package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
)

func registerRedirectRoutes(h *handlers.Handler, r *gin.Engine) {
	rGroup := r.Group("r")
	{
		rGroup.GET("/:shrturl", h.RedirectHandler)
	}
}
