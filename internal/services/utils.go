package services

import (
	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr/urlshrt/internal/pkg/hash"
	"github.com/wrtgvr2/errsuit"
)

func convertUserReqToUserDb(userReq *models_http.UserRequest) (*models_db.User, *errsuit.AppError) {
	hashedPassword, err := hash.HashPassword(userReq.Password)
	if err != nil {
		return nil, errsuit.NewInternal("internal server error", err, true)
	}

	userData := models_db.User{
		ID:           0,
		Username:     userReq.Username,
		PasswordHash: hashedPassword,
	}

	if userReq.DisplayUsername != nil {
		userData.DisplayUsername = userReq.Username
	}

	return &userData, nil
}
