package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
)

func DecodeBody(c *gin.Context, obj any) *errsuit.AppError {
	err := c.ShouldBindJSON(&obj)
	if err != nil {
		return errsuit.NewBadRequest("invalid body", err, false)
	}
	return nil
}
