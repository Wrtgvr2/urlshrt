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
	ulrModel := models_db.URL{
		UserID:    userId,
		OrigURL:   urlReq.URL,
		ShortURL:  shortUrl,
		ExpiresAt: &expirationDate,
	}

	createdUrl, err := s.Repo.CreateNewShortUrl(&ulrModel)
	if err != nil {
		return nil, err
	}

	return convertUrlDbToUrlResp(createdUrl), nil
}
