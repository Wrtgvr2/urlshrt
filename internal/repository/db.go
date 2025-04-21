package repository

import (
	"fmt"
	"log"
	"os"

	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models_db.User{},
		&models_db.URL{},
		&models_db.RefreshToken{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
