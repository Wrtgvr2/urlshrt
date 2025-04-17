package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/internal/handlers"
)

func registerShortUrls(h *handlers.Handler, r *gin.Engine) {
	r.GET("/r/:shrturl", h.ShortUrlsHandler)
}
