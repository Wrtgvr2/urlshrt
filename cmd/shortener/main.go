package main

import (
	"github.com/wrtgvr/urlshrt/internal/app"
	"github.com/wrtgvr/urlshrt/internal/router"
)

func main() {
	App := app.InitApp()
	r := router.RegisterRoutes(&App.Handler)

	r.Run(":8080")
}
