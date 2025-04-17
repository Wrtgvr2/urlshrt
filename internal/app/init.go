package app

import (
	"log"

	env "github.com/joho/godotenv"
	"github.com/wrtgvr/urlshrt/internal/handlers"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr/urlshrt/internal/services"
)

type App struct {
	Handler handlers.Handler
}

func InitApp() *App {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	rep.InitDatabase()

	h := handlers.Handler{
		UserServices: services.UserServices{
			Repo: rep.PostgresUserRepo{},
		},
		UrlServices: services.UrlServices{
			Repo: rep.PostgresUlrRepo{},
		}, // TODO: Make a constructor functions for handler and services
	}

	return &App{
		Handler: h,
	}
}
