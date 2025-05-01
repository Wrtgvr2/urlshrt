package services

import (
	"time"

	models_db "github.com/wrtgvr/urlshrt/internal/models/db"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr/urlshrt/internal/validation"
	"github.com/wrtgvr/urlshrt/pkg/hash"
	"github.com/wrtgvr/urlshrt/pkg/jwt"
	"github.com/wrtgvr2/errsuit"
)

func convertUserDbToUserResp(userDb *models_db.User) *models_http.UserResponse {
	return &models_http.UserResponse{
		ID:              userDb.ID,
		Username:        userDb.Username,
		DisplayUsername: userDb.DisplayUsername,
	}
}

func convertUserReqToUserDb(userReq *models_http.UserRequest) (*models_db.User, *errsuit.AppError) {
	hashedPassword, err := hash.HashPassword(userReq.Password)
	if err != nil {
		return nil, errsuit.NewInternal("internal server error", err, true)
	}

	userData := models_db.User{
		Username:     userReq.Username,
		PasswordHash: hashedPassword,
	}

	if userReq.DisplayUsername != nil {
		userData.DisplayUsername = userReq.Username
	} else {
		userData.DisplayUsername = userData.Username
	}

	return &userData, nil
}

func validateUserData(username, password string) *errsuit.AppError {
	err := validation.ValidateUsername(username)
	if err != nil {
		return err
	}
	err = validation.ValidatePassword(password)
	if err != nil {
		return err
	}
	return nil
}

func createRefreshTokenModel(tokenStr string) (*models_db.RefreshToken, *errsuit.AppError) {
	tokenJti, err := jwt.GetJtiFromRefreshToken(tokenStr)
	if err != nil {
		return nil, errsuit.NewInternal("unable to get jti from token", err, true)
	}
	tokenUserId, err := jwt.GetUserIdFromToken(tokenStr)
	if err != nil {
		return nil, errsuit.NewInternal("unable to get user ID from token", err, true)
	}
	tokenExp, err := jwt.GetTokenExpirationUnixTime(tokenStr)
	if err != nil {
		return nil, errsuit.NewInternal("unable to get expiration time from token", err, true)
	}
	return &models_db.RefreshToken{
		JTI:       tokenJti,
		UserID:    tokenUserId,
		ExpiresAt: time.Unix(tokenExp, 0),
	}, nil
}
