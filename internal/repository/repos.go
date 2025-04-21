package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
)

type UserRepo interface {
	GetUserByUsername(string) (*models_db.User, *errsuit.AppError)
}

type UrlRepo interface{}

type RefreshTokenRepo interface {
	CreateRefreshTokenInfo(*models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError)
}
