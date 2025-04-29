package services

import (
	"github.com/google/uuid"
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr2/errsuit"
)

type TokenServices struct {
	Repo rep.TokenRepo
}

func NewTokenServices(repo rep.TokenRepo) TokenServices {
	return TokenServices{
		Repo: repo,
	}
}

func (s *TokenServices) UpdateRefreshToken(jti uuid.UUID, newTokenStr string) (*models_db.RefreshToken, *errsuit.AppError) {
	_, appErr := s.Repo.GetNotRevokedTokenByJTI(jti)
	if appErr != nil {
		return nil, appErr
	}

	appErr = s.Repo.RevokeToken(jti)
	if appErr != nil {
		return nil, appErr
	}

	tokenModel, appErr := createRefreshTokenModel(newTokenStr)
	if appErr != nil {
		return nil, appErr
	}

	newToken, appErr := s.Repo.CreateRefreshTokenInfo(tokenModel)
	if appErr != nil {
		return nil, appErr
	}
	return newToken, nil
}
