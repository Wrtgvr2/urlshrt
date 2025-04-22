package app

import (
	"log"

	env "github.com/joho/godotenv"
	"github.com/wrtgvr/urlshrt/api/handlers"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr/urlshrt/internal/services"
	"gorm.io/gorm"
)

type App struct {
	Handler *handlers.Handler
}

func InitApp() *App {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	db := rep.InitDatabase()
	h := initHandler(db)

	return &App{
		Handler: h,
	}
}

func initHandler(db *gorm.DB) *handlers.Handler {
	userRepo := rep.NewPostgresUserRepo(db)
	urlRepo := rep.NewPostgresUrlRepo(db)

	userServices := services.NewUserServices(userRepo)
	urlServices := services.NewUrlServices(urlRepo)

	h := handlers.NewHandler(&userServices, &urlServices)

	return h
}
