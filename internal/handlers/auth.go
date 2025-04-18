package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	//user, err := h.UserServices.Login(&req)
}
