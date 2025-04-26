package repository

import (
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
	GetTokenByJTI(jti string) (*models_db.RefreshToken, *errsuit.AppError)
	GetNotRevokedTokenByJTI(jti string) (*models_db.RefreshToken, *errsuit.AppError)
	RevokeToken(jti string) *errsuit.AppError
	ReplaceRefreshToken(oldTokenJTI string, newTokenData models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
}
