package repository

import (
	"fmt"

	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
	"gorm.io/gorm"
)

type PostgresUrlRepo struct {
	DB *gorm.DB
}

func NewPostgresUrlRepo(db *gorm.DB) *PostgresUrlRepo {
	return &PostgresUrlRepo{DB: db}
}

func (p *PostgresUrlRepo) CreateNewShortUrl(urlModel *models_db.URL) (*models_db.URL, *errsuit.AppError) {
	res := p.DB.Create(urlModel)
	if res.Error != nil {
		return nil, errsuit.NewInternal("unable to create url", res.Error, true)
	}
	return urlModel, nil
}

func (p *PostgresUrlRepo) getUrl(searchField string, fieldValue any) (*models_db.URL, *errsuit.AppError) {
	var url models_db.URL
	res := p.DB.First(&url, fmt.Sprintf("%s = ?", searchField), fieldValue)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, errsuit.NewNotFound("url not found", res.Error, false)
		}
		return nil, errsuit.NewInternal("unable to get url", res.Error, true)
	}

	return &url, nil
}

func (p *PostgresUrlRepo) GetUrlByShortUrl(shortUrl string) (*models_db.URL, *errsuit.AppError) {
	url, err := p.getUrl("short_url", shortUrl)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (p *PostgresUrlRepo) GetValidUrlByShortUrl(shortUrl string) (*models_db.URL, *errsuit.AppError) {
	url, err := p.GetUrlByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}

	if url.Revoked {
		// Return "Not Found" error cuz there is no sense to say user "Hey, this link is revoked"
		// let user (not like user-user, user as guy who just click a link) think there is no url
		return nil, errsuit.NewNotFound("short url not found", err, false)
	}
	return url, nil
}

func (p *PostgresUrlRepo) IncrementRedirectCount(url *models_db.URL) *errsuit.AppError {
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

func (p *PostgresUrlRepo) GetUrlById(id uint64) (*models_db.URL, *errsuit.AppError) {
	url, err := p.getUrl("id", id)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (p *PostgresUrlRepo) GetUrlByUserId(userId uint64) (*models_db.URL, *errsuit.AppError) {
	url, err := p.getUrl("user_id", userId)
	if err != nil {
		return nil, err
	}
	return url, nil
}
