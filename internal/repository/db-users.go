package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo(db *gorm.DB) *PostgresUserRepo {
	return &PostgresUserRepo{DB: db}
}

func (p *PostgresUserRepo) GetUserByUsername(username string) (*models_db.User, *errsuit.AppError) {
	var user models_db.User
	res := p.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, errsuit.NewNotFound("user not found", res.Error, false)
		}
		return nil, errsuit.NewInternal("DB error", res.Error, true)
	}

	return &user, nil
}

func (p *PostgresUserRepo) CreateUser(userData *models_db.User) (*models_db.User, *errsuit.AppError) {
	user := *userData
	res := p.DB.Select("*").Create(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errsuit.NewInternal("unable to create user", res.Error, true)
	}

	return &user, nil
}
