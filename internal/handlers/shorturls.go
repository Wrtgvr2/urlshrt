package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortUrlsHandler(c *gin.Context) {
	shortUrl := c.Param("shrturl")

	url := shortUrl
	// search for shortUrl in DB if shortUrl exists then get actual url and redirect
	c.Redirect(http.StatusFound, url)
}
