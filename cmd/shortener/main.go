package main

import (
	"github.com/wrtgvr/urlshrt/internal/router"
)

func main() {
	r := router.RegisterRoutes()

	r.Run(":8080")
}
