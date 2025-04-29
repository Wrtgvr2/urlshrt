package repository

import (
	"github.com/google/uuid"
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

func (p *PostgresTokenRepo) GetTokenByJTI(jti uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError) {
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

func (p *PostgresTokenRepo) GetNotRevokedTokenByJTI(jti uuid.UUID) (*models_db.RefreshToken, *errsuit.AppError) {
	token, err := p.GetTokenByJTI(jti)
	if err != nil {
		return nil, err
	}
	if token.Revoked {
		return nil, errsuit.NewUnauthorized("invalid token", err, false)
	}
	return token, nil
}

func (p *PostgresTokenRepo) RevokeToken(jti uuid.UUID) *errsuit.AppError {
	token, appErr := p.GetTokenByJTI(jti)
	if appErr != nil {
		return appErr
	}
	token.Revoked = true
	err := p.DB.Save(token).Error
	if err != nil {
		return errsuit.NewInternal("unable to update token info", err, true)
	}
	return nil
}

func (p *PostgresTokenRepo) ReplaceRefreshToken(oldTokenJTI uuid.UUID, newTokenData models_db.RefreshToken) (*models_db.RefreshToken, *errsuit.AppError) {
	err := p.RevokeToken(oldTokenJTI)
	if err != nil {
		return nil, err
	}
	newToken, err := p.CreateRefreshTokenInfo(&newTokenData)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
