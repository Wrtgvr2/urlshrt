package services

import (
	"errors"
	"net/http"

	"github.com/wrtgvr/urlshrt/internal/apperrors"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr/urlshrt/internal/pkg/hash"
	"github.com/wrtgvr/urlshrt/internal/pkg/jwt"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
)

type UserServices struct {
	Repo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{Repo: repo}
}

func (s *UserServices) Login(userReq models_http.UserRequest) (string, string, *apperrors.AppError) {
	user, appErr := s.Repo.GetUserByUsername(userReq.Username)
	if appErr != nil {
		return "", "", appErr
	}
	if passwordCorrect := hash.CheckPasswordHash(userReq.Password, user.PasswordHash); !passwordCorrect {
		return "", "", apperrors.WrapError(errors.New("password mismatch"), http.StatusUnauthorized, "password mismatch")
	}
	accessToken, refreshToken, err := jwt.CreateTokens(user.ID)
	if err != nil {
		return "", "", apperrors.WrapError(err, http.StatusInternalServerError, "internal server error")
	}
	return accessToken, refreshToken, nil
}
