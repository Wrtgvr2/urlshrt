package repository

import (
	"fmt"

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

func (p *PostgresUserRepo) getUser(searchField string, fieldValue any) (*models_db.User, *errsuit.AppError) {
	var user models_db.User
	res := p.DB.First(&user, fmt.Sprintf("%s = ?", searchField), fieldValue)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, errsuit.NewNotFound("user not found", res.Error, false)
		}
		return nil, errsuit.NewInternal("DB error", res.Error, true)
	}

	return &user, nil
}

func (p *PostgresUserRepo) GetUserById(id uint64) (*models_db.User, *errsuit.AppError) {
	return p.getUser("id", id)
}

func (p *PostgresUserRepo) GetUserByUsername(username string) (*models_db.User, *errsuit.AppError) {
	return p.getUser("username", username)
}

func (p *PostgresUserRepo) GetAllUsers() ([]models_db.User, *errsuit.AppError) {
	var users []models_db.User
	err := p.DB.Find(&users).Error
	if err != nil {
		return nil, errsuit.NewInternal("can't get all users", err, true)
	}
	return users, nil
}

func (p *PostgresUserRepo) CreateUser(userData *models_db.User) (*models_db.User, *errsuit.AppError) {
	user := *userData
	res := p.DB.Select("*").Create(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errsuit.NewInternal("unable to create user", res.Error, true)
	}

	return &user, nil
}

func (p *PostgresUserRepo) DeleteUser(id uint64) *errsuit.AppError {
	err := p.DB.Where("id = ?", id).Delete(&models_db.User{}).Error
	if err != nil {
		return errsuit.NewInternal("can't delete user", err, true)
	}

	return nil
}
