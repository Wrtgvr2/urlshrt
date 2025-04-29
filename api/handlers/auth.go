package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr/urlshrt/pkg/jwt"
	"github.com/wrtgvr2/errsuit"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	req := models_http.UserRequest{}
	appErr := DecodeBody(c, &req)
	if ginadap.HandleError(c, appErr) {
		return
	}

	accessToken, refreshToken, err := h.AuthServices.Login(&req)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) RegisterHandler(c *gin.Context) {
	req := models_http.UserRequest{}
	appErr := DecodeBody(c, &req)
	if ginadap.HandleError(c, appErr) {
		return
	}

	user, err := h.AuthServices.Register(&req)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) RefreshTokenHandler(c *gin.Context) {
	jtiAny, exists := c.Get("JTI")
	if !exists {
		ginadap.HandleError(c, errsuit.NewUnauthorized("jti not found in context", nil, false))
		return
	}

	jti, ok := jtiAny.(uuid.UUID)
	if !ok {
		ginadap.HandleError(c, errsuit.NewUnauthorized("invalid jti", nil, false))
		return
	}

	userIdAny, exists := c.Get("UserID")
	if !exists {
		ginadap.HandleError(c, errsuit.NewUnauthorized("user id not found in context", nil, false))
		return
	}
	userId, ok := userIdAny.(uint64)
	if !ok {
		ginadap.HandleError(c, errsuit.NewInternal("invalid user id type in context", nil, true))
		return
	}

	newAccessToken, newRefreshToken, err := jwt.CreateTokens(userId)
	if ginadap.HandleError(c, err) {
		return
	}

	_, err = h.TokenServices.UpdateRefreshToken(jti, newRefreshToken)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
