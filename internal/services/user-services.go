package services

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr/urlshrt/pkg/hash"
	"github.com/wrtgvr/urlshrt/pkg/jwt"
	"github.com/wrtgvr2/errsuit"
)

type UserServices struct {
	Repo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{Repo: repo}
}

func (s *UserServices) Login(userReq *models_http.UserRequest) (string, string, *errsuit.AppError) {
	verr := ValidateUserData(userReq.Username, userReq.Password)
	if verr != nil {
		return "", "", verr
	}
	user, appErr := s.Repo.GetUserByUsername(userReq.Username)
	if appErr != nil {
		return "", "", appErr
	}
	if passwordCorrect := hash.CheckPasswordWithHash(userReq.Password, user.PasswordHash); !passwordCorrect {
		return "", "", errsuit.NewUnauthorized("password mismatch", nil, false)
	}
	accessToken, refreshToken, err := jwt.CreateTokens(user.ID)
	if err != nil {
		return "", "", errsuit.NewInternal("failed to create tokens", err, true)
	}
	return accessToken, refreshToken, nil
}

func (s *UserServices) Register(userReq *models_http.UserRequest) (*models_db.User, *errsuit.AppError) {
	err := ValidateUserData(userReq.Username, userReq.Password)
	if err != nil {
		return nil, err
	}
	user, appErr := s.Repo.GetUserByUsername(userReq.Username)
	if appErr != nil && appErr.Type != errsuit.TypeNotFound {
		return nil, appErr
	}
	if user != nil {
		return nil, errsuit.NewConflict("username is already taken", nil, false)
	}

	userData, appErr := convertUserReqToUserDb(userReq)
	if appErr != nil {
		return nil, appErr
	}

	createdUser, appErr := s.Repo.CreateUser(userData)
	if appErr != nil {
		return nil, appErr
	}

	return createdUser, nil
}
