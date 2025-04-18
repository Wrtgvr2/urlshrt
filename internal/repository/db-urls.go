package repository

import (
	"gorm.io/gorm"
)

type PostgresUlrRepo struct {
	DB *gorm.DB
}

func NewPostgresUrlRepo(db *gorm.DB) *PostgresUlrRepo {
	return &PostgresUlrRepo{DB: db}
}
