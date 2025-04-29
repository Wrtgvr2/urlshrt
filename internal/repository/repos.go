package repository

import (
	"github.com/google/uuid"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
)

type UserRepo interface {
	GetUserById(id uint64) (*models_db.User, *errsuit.AppError)
	GetUserByUsername(string) (*models_db.User, *errsuit.AppError)
	CreateUser(*models_db.User) (*models_db.User, *errsuit.AppError)
}

type UrlRepo interface{}

type TokenRepo interface {
	CreateRefreshTokenInfo(tokenData *models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
	GetTokenByJTI(jti uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError)
	GetNotRevokedTokenByJTI(jti uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError)
	RevokeToken(jti uuid.UUID) *errsuit.AppError
	ReplaceRefreshToken(oldTokenJTI uuid.UUID, newTokenData models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
}
