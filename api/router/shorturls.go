package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/internal/handlers"
)

func registerShortUrls(h *handlers.Handler, r *gin.Engine) {
	group := r.Group("r")
	{
		group.GET("/:shrturl", h.RedirectHandler)
		//group.POST("/shorten", h.ShortenHandler)
	}
}
