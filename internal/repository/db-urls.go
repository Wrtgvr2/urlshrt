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

func (p *PostgresUlrRepo) GetUrlByShortUrl(shortUrl string) (*models_db.URL, *errsuit.AppError) {
	var urlModel models_db.URL
	err := p.DB.Where("short_url = ?", shortUrl).First(&urlModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errsuit.NewNotFound("short url not found", err, false)
		}
		return nil, errsuit.NewInternal("unable to get orig url", err, true)
	}

	return &urlModel, nil
}

func (p *PostgresUlrRepo) GetValidUrlByShortUrl(shortUrl string) (*models_db.URL, *errsuit.AppError) {
	url, err := p.GetUrlByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}

	if url.Revoked {
		// Return "Not Found" error cuz there is no sense to say user "Hey, this link is revoked" let user think there is no url
		return nil, errsuit.NewNotFound("short url not found", err, false)
	}
	return url, nil
}

func (p *PostgresUlrRepo) IncrementRedirectCount(url *models_db.URL) *errsuit.AppError {
	url.Redirects += 1
	res := p.DB.Save(url)
	if res.Error != nil {
		return errsuit.NewInternal("unable to increment redirect count", res.Error, true)
	}
	if res.RowsAffected == 0 {
		return errsuit.NewInternal("unable to increment redirect count", res.Error, true)
	}

	return nil
}
