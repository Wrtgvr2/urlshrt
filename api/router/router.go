package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/internal/handlers"
)

func RegisterRoutes(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	registerShortUrls(h, r)

	return r
}
