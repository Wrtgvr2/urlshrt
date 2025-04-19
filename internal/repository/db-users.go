package repository

import (
	"net/http"

	"github.com/wrtgvr/urlshrt/internal/apperrors"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{DB: db}
}

func (p *PostgresUserRepo) GetUserByUsername(username string) (*models_db.User, *apperrors.AppError) {
	var user models_db.User
	res := p.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, apperrors.WrapError(res.Error, http.StatusNotFound, "user not found")
		}
		return nil, apperrors.WrapError(res.Error, http.StatusInternalServerError, "internal server error")
	}

	return &user, nil
}
