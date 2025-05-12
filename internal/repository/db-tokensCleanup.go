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

func (p *PostgresTokensCleanupRepo) DeleteTooOldTokens() (int64, error) {
	res := p.DB.Where("expires_at < ?", time.Now().Add(-30*24*time.Hour)).Delete(&models_db.RefreshToken{})
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *PostgresTokensCleanupRepo) RevokeExpiredTokens() (int64, error) {
	res := p.DB.Model(&models_db.URL{}).Where("expires_at < ?", time.Now()).Update("revoked", true)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
