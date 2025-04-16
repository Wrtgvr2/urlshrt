package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	registerShortUrls(r)

	return r
}
