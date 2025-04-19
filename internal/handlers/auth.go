package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	req := models_http.UserRequest{}
	appErr := DecodeBody(c, req)
	if HandleError(c, appErr) {
		return
	}

	accessToken, refreshToken, err := h.UserServices.Login(req)
	if HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
