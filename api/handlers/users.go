package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) GetUserHandler(c *gin.Context) {
	id, err := GetIdFromContext(c)
	if ginadap.HandleError(c, err) {
		return
	}

	user, appErr := h.UserServices.GetUser(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, user)
}

func (h *Handler) GetUsersHandler(c *gin.Context) {
	user, appErr := h.UserServices.GetUsers()
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, user)
}
