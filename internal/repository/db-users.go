package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{DB: db}
}

func (p *PostgresUserRepo) GetUserByUsername(username string) (*models_db.User, error) {
	return nil, nil
}
