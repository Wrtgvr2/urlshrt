package handlers

import (
	"github.com/gin-gonic/gin"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) GetUserHandler(c *gin.Context) {
	id, err := GetIdFromContextParam(c)
	if ginadap.HandleError(c, err) {
		return
	}

	user, appErr := h.UserServices.GetUser(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, user)
}

func (h *Handler) DeleteUserHandler(c *gin.Context) {
	id, err := GetIdFromContextParam(c)
	if ginadap.HandleError(c, err) {
		return
	}

	err = h.UserServices.DeleteUser(id)
	if ginadap.HandleError(c, err) {
		return
	}

	c.Status(204)
}

func (h *Handler) PatchUserHandler(c *gin.Context) {
	id, err := GetIdFromContextParam(c)
	if ginadap.HandleError(c, err) {
		return
	}

	var userReq models_http.UserPatchRequest
	err = DecodeBody(c, &userReq)
	if ginadap.HandleError(c, err) {
		return
	}

	updatedUser, err := h.UserServices.PatchUser(id, &userReq)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(200, updatedUser)
}
