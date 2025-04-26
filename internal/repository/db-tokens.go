package repository

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	"github.com/wrtgvr2/errsuit"
	"gorm.io/gorm"
)

type PostgresTokenRepo struct {
	DB *gorm.DB
}

func NewPostgresTokenRepo(db *gorm.DB) *PostgresTokenRepo {
	return &PostgresTokenRepo{db}
}

func (p *PostgresTokenRepo) CreateRefreshTokenInfo(tokenData *models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError) {
	res := p.DB.Create(tokenData)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errsuit.NewInternal("unable to create token info", res.Error, true)
	}

	return tokenData, nil
}

func (p *PostgresTokenRepo) GetTokenByJTI(jti string) (*models_db.RefreshToken, *errsuit.AppError) {
	var token models_db.RefreshToken
	err := p.DB.First(&token, jti).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errsuit.NewNotFound("token not found", err, false)
		}
		return nil, errsuit.NewInternal("internal server error", err, true)
	}
	return &token, nil
}

func (p *PostgresTokenRepo) GetNotRevokedTokenByJTI(jti string) (*models_db.RefreshToken, *errsuit.AppError) {
	token, err := p.GetTokenByJTI(jti)
	if err != nil {
		return nil, err
	}
	if token.Revoked {
		return nil, errsuit.NewUnauthorized("invalid token", err, false)
	}
	return token, nil
}

func (p *PostgresTokenRepo) ReplaceRefreshToken(oldTokenJTI, tokenStr string) (*models_db.RefreshToken, *errsuit.AppError) {
	oldToken, appErr := p.GetNotRevokedTokenByJTI(oldTokenJTI)
	if appErr != nil {
		return nil, appErr
	}
	oldToken.Revoked = true
	err := p.DB.Save(oldToken).Error
	if err != nil {
		return nil, errsuit.NewInternal("unable to update token info", err, true)
	}

	//newToken, err := p.CreateRefreshTokenInfo()

	// При входе - создавать токены
	// При обновлении токенов - отзывать прошлый токен, создавать новый, с новым JTI
	// TODO:
	// 1. Создание токена в БД при создании токенов при логине
	// 2. Завершить функцию "замены" токенов (удаление прошлого, создание нового)
	// 3. ???
}
