package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
)

func DecodeBody(c *gin.Context, obj any) *errsuit.AppError {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return errsuit.NewBadRequest("invalid body", err, false)
	}
	return nil
}

func GetIdFromContextParam(c *gin.Context) (uint64, *errsuit.AppError) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, errsuit.NewBadRequest("invalid id", err, false)
	}

	return id, nil
}

func GetUserIdFromContext(c *gin.Context) (uint64, *errsuit.AppError) {
	idAny, exists := c.Get("UserID")
	if !exists {
		return 0, errsuit.NewBadRequest("invalid id", nil, false)
	}
	id, ok := idAny.(uint64)
	if !ok {
		return 0, errsuit.NewBadRequest("invalid id", nil, false)
	}

	return id, nil
}
