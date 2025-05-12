package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) GetUrlHandler(c *gin.Context) {
	id, appErr := GetIdFromContextParam(c)
	if ginadap.HandleError(c, appErr) {
		return
	}

	url, appErr := h.UrlServices.GetUrlById(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, url)
}

func (h *Handler) GetUserUrlsHandler(c *gin.Context) {
	id, appErr := GetUserIdFromContext(c)
	if ginadap.HandleError(c, appErr) {
		return
	}

	url, appErr := h.UrlServices.GetUserUrls(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.JSON(200, url)
}

func (h *Handler) DeleteUrlHandler(c *gin.Context) {
	id, appErr := GetIdFromContextParam(c)
	if ginadap.HandleError(c, appErr) {
		return
	}

	appErr = h.UrlServices.DeleteUrl(id)
	if ginadap.HandleError(c, appErr) {
		return
	}

	c.Status(204)
}
