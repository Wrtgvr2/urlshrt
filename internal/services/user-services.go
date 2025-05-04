package services

import (
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	rep "github.com/wrtgvr/urlshrt/internal/repository"
	"github.com/wrtgvr/urlshrt/pkg/hash"
	"github.com/wrtgvr2/errsuit"
)

type UserServices struct {
	UserRepo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{repo}
}

func (s UserServices) GetUser(id uint64) (*models_http.UserResponse, *errsuit.AppError) {
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

func (s *UserServices) DeleteUser(id uint64) *errsuit.AppError {
	err := s.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServices) PatchUser(id uint64, userReq *models_http.UserPatchRequest) (*models_http.UserResponse, *errsuit.AppError) {
	existingUser, err := s.UserRepo.GetUserByUsername(*userReq.Username)
	if err != nil && err.Type != errsuit.TypeNotFound {
		return nil, err
	}
	if existingUser != nil {
		if existingUser.ID != id {
			return nil, errsuit.NewConflict("username already taken", nil, false)
		}
	}

	user, err := s.UserRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if userReq.Username != nil {
		user.Username = *userReq.Username
	}
	if userReq.DisplayUsername != nil {
		user.DisplayUsername = *userReq.DisplayUsername
	}
	if userReq.Password != nil {
		hashedPassword, err := hash.HashPassword(*userReq.Password)
		if err != nil {
			return nil, errsuit.AsAppError(err)
		}
		user.PasswordHash = hashedPassword
	}

	updatedUser, err := s.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	resp := convertUserDbToUserResp(updatedUser)
	return resp, nil
}
