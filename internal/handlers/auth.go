package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	req := models_http.UserRequest{}
	appErr := DecodeBody(c, req)
	if ginadap.HandleError(c, appErr) {
		return
	}

	accessToken, refreshToken, err := h.UserServices.Login(req)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
