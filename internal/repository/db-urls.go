package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
	"gorm.io/gorm"
)

type PostgresUlrRepo struct {
	DB *gorm.DB
}

func NewPostgresUrlRepo(db *gorm.DB) *PostgresUlrRepo {
	return &PostgresUlrRepo{DB: db}
}

func (p *PostgresUlrRepo) CreateNewShortUrl(urlModel *models_db.URL) (*models_db.URL, *errsuit.AppError) {
	res := p.DB.Create(urlModel)
	if res.Error != nil {
		return nil, errsuit.NewInternal("unable to create url", res.Error, true)
	}
	return urlModel, nil
}
