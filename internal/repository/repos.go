package repository

import (
	"github.com/wrtgvr/urlshrt/internal/apperrors"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
)

type UserRepo interface {
	GetUserByUsername(string) (*models_db.User, *apperrors.AppError)
}

type UrlRepo interface{}
