package main

import (
	"github.com/wrtgvr/urlshrt/api/router"
	"github.com/wrtgvr/urlshrt/internal/app"
)

func main() {
	App := app.InitApp()
	r := router.RegisterRoutes(App.Handler)

	r.Run(":8080")
}
