package services

import (
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
)

type UserServices struct {
	UserRepo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{repo}
}

func (s UserServices) GetUser(id uint64) (*models_http.UserResponse, error) {
	user, appErr := s.UserRepo.GetUserById(id)
	if appErr != nil {
		return nil, appErr
	}

	userRes := models_http.UserResponse{
		ID:              id,
		Username:        user.Username,
		DisplayUsername: user.DisplayUsername,
	}

	return &userRes, nil
}
