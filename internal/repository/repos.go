package repository

import (
	"github.com/google/uuid"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
)

type UserRepo interface {
	GetUserById(uint64) (*models_db.User, *errsuit.AppError)
	GetUserByUsername(string) (*models_db.User, *errsuit.AppError)
	CreateUser(*models_db.User) (*models_db.User, *errsuit.AppError)
	DeleteUser(uint64) *errsuit.AppError
	UpdateUser(*models_db.User) (*models_db.User, *errsuit.AppError)
}

type UrlRepo interface {
	CreateNewShortUrl(urlModel *models_db.URL) (*models_db.URL, *errsuit.AppError)
	GetUrlByShortUrl(string) (*models_db.URL, *errsuit.AppError)
	GetValidUrlByShortUrl(string) (*models_db.URL, *errsuit.AppError)
	IncrementRedirectCount(*models_db.URL) *errsuit.AppError
}

type TokenRepo interface {
	CreateRefreshTokenInfo(*models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
	GetTokenByJTI(uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError)
	RevokeToken(uuid.UUID) *errsuit.AppError
	ReplaceRefreshToken(uuid.UUID, models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
	GetValidTokenByJti(uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError)
}

type TokensCleanupRepo interface {
	DeleteTooOldTokens() (int64, error)
}
