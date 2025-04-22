package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
)

func RegisterRoutes(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	registerShortUrlsRouter(h, r)
	registerUserRouter(h, r)

	return r
}
