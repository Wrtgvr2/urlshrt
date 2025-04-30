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

func LoadEnvVars() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}
}

func InitApp() *App {
	LoadEnvVars()
	db := rep.InitDatabase()
	h := initHandler(db)

	return &App{
		Handler: h,
	}
}

func initHandler(db *gorm.DB) *handlers.Handler {
	userRepo := rep.NewPostgresUserRepo(db)
	urlRepo := rep.NewPostgresUrlRepo(db)
	tokenRepo := rep.NewPostgresTokenRepo(db)

	userServices := services.NewUserServices(userRepo)
	urlServices := services.NewUrlServices(urlRepo)
	authServices := services.NewAuthServices(userRepo, tokenRepo)
	tokenServices := services.NewTokenServices(tokenRepo)

	h := handlers.NewHandler(
		&userServices, &urlServices,
		&authServices, &tokenServices,
	)

	return h
}
