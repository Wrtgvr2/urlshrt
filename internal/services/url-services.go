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
	salt := urlReq.URL + fmt.Sprintf("%d%d%s", userId, time.Now().UnixNano(), os.Getenv("URL_SECRET_SALT"))
	hash := sha1.Sum([]byte(salt))
	shortUrl := base62.EncodeToString(hash[:8])

	expirationDate := time.Now().Add(30 * 24 * time.Hour)
	urlModel := models_db.URL{
		UserID:    userId,
		OrigURL:   urlReq.URL,
		ShortURL:  shortUrl,
		ExpiresAt: &expirationDate,
	}

	createdUrl, appErr := s.Repo.CreateNewShortUrl(&urlModel)
	if appErr != nil {
		return nil, appErr
	}

	return convertUrlDbToUrlResp(createdUrl), nil
}

func (s *UrlServices) GetValidUrlByShortUrl(url string) (*models_db.URL, *errsuit.AppError) {
	origUrl, appErr := s.Repo.GetValidUrlByShortUrl(url)
	if appErr != nil {
		return nil, appErr
	}

	return origUrl, nil
}

func (s *UrlServices) GetUrlById(id uint64) (*models_http.UrlResponse, *errsuit.AppError) {
	url, appErr := s.Repo.GetUrlById(id)
	if appErr != nil {
		return nil, appErr
	}

	urlResp := convertUrlDbToUrlResp(url)

	return urlResp, nil
}

func (s *UrlServices) GetUrlByUserId(id uint64) (*models_http.UrlResponse, *errsuit.AppError) {
	url, appErr := s.Repo.GetUrlByUserId(id)
	if appErr != nil {
		return nil, appErr
	}
	urlResp := convertUrlDbToUrlResp(url)

	return urlResp, nil
}

func (s *UrlServices) IncrementRedirectCount(url *models_db.URL) *errsuit.AppError {
	err := s.Repo.IncrementRedirectCount(url)
	return err
}

// Accep user id and url id
func (s *UrlServices) DeleteUrl(userId, urlId uint64) *errsuit.AppError {
	url, appErr := s.Repo.GetUrlById(urlId)
	if appErr != nil {
		return appErr
	}
	if url.UserID != userId {
		return errsuit.NewForbidden("you do not have permission to perform this action", nil, true)
	}

	err := s.Repo.DeleteUrl(urlId)
	return err
}

func (s *UrlServices) GetUserUrls(userId uint64) ([]models_http.UrlResponse, *errsuit.AppError) {
	urls, appErr := s.Repo.GetUserUrls(userId)
	if appErr != nil {
		return nil, appErr
	}

	urlsResp := []models_http.UrlResponse{}
	for _, url := range urls {
		urlsResp = append(urlsResp, *convertUrlDbToUrlResp(&url))
	}

	return urlsResp, nil
}
