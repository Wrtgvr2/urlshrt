package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, code int) bool {
	if err == nil {
		return false
	}

	msg := err.Error()

	if code == http.StatusInternalServerError {
		msg = http.StatusText(http.StatusInternalServerError)
	}

	c.JSON(code, gin.H{
		"error": msg,
	})

	return true
}
