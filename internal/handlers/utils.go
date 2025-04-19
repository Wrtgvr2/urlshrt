package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/urlshrt/internal/apperrors"
)

func HandleError(c *gin.Context, appErr *apperrors.AppError) bool {
	if appErr == nil {
		return false
	}

	msg := appErr.Message

	if appErr.StatusCode == http.StatusInternalServerError {
		msg = http.StatusText(http.StatusInternalServerError)
	}

	c.JSON(appErr.StatusCode, gin.H{
		"error": msg,
	})

	return true
}

func DecodeBody(c *gin.Context, obj any) *apperrors.AppError {
	err := c.ShouldBindJSON(&obj)
	if err != nil {
		return apperrors.WrapError(err, http.StatusBadRequest, "invalid credentials")
	}
	return nil
}
