package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
	"gorm.io/gorm"
)

type PostgresRefreshToken struct {
	DB *gorm.DB
}

func (p *PostgresRefreshToken) CreateRefreshTokenInfo(tokenData *models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError) {
	res := p.DB.Create(tokenData)
	if res.Error != nil {
		return nil, errsuit.NewInternal("DB error:", res.Error, true)
	}
	if res.RowsAffected == 0 {
		return nil, errsuit.NewInternal("DB error:", res.Error, true)
	}

	return tokenData, nil
}
