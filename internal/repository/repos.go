package repository

import models_db "github.com/wrtgvr/urlshrt/internal/models/db"

type UserRepo interface {
	GetUserByUsername(string) (*models_db.User, error)
}

type UrlRepo interface{}
