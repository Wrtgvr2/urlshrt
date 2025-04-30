package repository

import (
	"time"

	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"gorm.io/gorm"
)

type PostgresTokensCleanupRepo struct {
	DB *gorm.DB
}

func NewPostgresTokensCleanupRepo(db *gorm.DB) *PostgresTokensCleanupRepo {
	return &PostgresTokensCleanupRepo{
		DB: db,
	}
}

func (r *PostgresTokensCleanupRepo) DeleteTooOldTokens() (int64, error) {
	res := r.DB.Where("expires_at < ?", time.Now().Add(-30*24*time.Hour)).Delete(&models_db.RefreshToken{})
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
