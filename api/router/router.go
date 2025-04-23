package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/api/handlers"
)

func RegisterRoutes(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	registerApiRoutes(h, r)
	registerAuthRouter(h, r)

	return r
}
