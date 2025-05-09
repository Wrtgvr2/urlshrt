package services

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/jxskiss/base62"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr2/errsuit"
)

type UrlServices struct {
	Repo rep.UrlRepo
}

func NewUrlServices(repo rep.UrlRepo) UrlServices {
	return UrlServices{Repo: repo}
}

func (s *UrlServices) CreateNewShortUrl(userId uint64, urlReq models_http.UrlRequest) (*models_http.UrlResponse, *errsuit.AppError) {
	salt := urlReq.URL + fmt.Sprintf("%d%d%s", userId, time.Now().UnixNano(), os.Getenv("URL_ULTRA_SECRET_EXTRA_SALT"))
	hash := sha1.Sum([]byte(salt))
	shortUrl := base62.EncodeToString(hash[:8])

	expirationDate := time.Now().Add(30 * 24 * time.Hour)
	urlModel := models_db.URL{
		UserID:    userId,
		OrigURL:   urlReq.URL,
		ShortURL:  shortUrl,
		ExpiresAt: &expirationDate,
	}

	createdUrl, err := s.Repo.CreateNewShortUrl(&urlModel)
	if err != nil {
		return nil, err
	}

	return convertUrlDbToUrlResp(createdUrl), nil
}

func (s *UrlServices) GetValidUrlByShortUrl(url string) (*models_db.URL, *errsuit.AppError) {
	origUrl, err := s.Repo.GetValidUrlByShortUrl(url)
	if err != nil {
		return nil, err
	}

	return origUrl, nil
}

func (s *UrlServices) GetUrlById(id uint64) (*models_http.UrlResponse, *errsuit.AppError) {
	url, err := s.Repo.GetUrlById(id)
	if err != nil {
		return nil, err
	}

	urlResp := convertUrlDbToUrlResp(url)

	return urlResp, nil
}

func (s *UrlServices) IncrementRedirectCount(url *models_db.URL) *errsuit.AppError {
	err := s.Repo.IncrementRedirectCount(url)
	return err
}
