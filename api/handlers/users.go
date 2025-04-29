package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) GetUserHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ginadap.HandleError(c, errsuit.NewBadRequest("invalid id", err, false))
		return
	}

	user, appErr := h.UserServices.GetUser(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, user)
}
