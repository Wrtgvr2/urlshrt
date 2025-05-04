package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models_http "github.com/wrtgvr/urlshrt/internal/models/http"
	"github.com/wrtgvr2/errsuit/drivers/ginadap"
)

func (h *Handler) RedirectHandler(c *gin.Context) {
	shortUrl := c.Param("shrturl")

	url, err := h.UrlServices.GetValidUrlByShortUrl(shortUrl)
	if ginadap.HandleError(c, err) {
		return
	}

	c.Redirect(http.StatusFound, url.OrigURL)

	err = h.UrlServices.IncrementRedirectCount(url)
}

func (h *Handler) ShortenHandler(c *gin.Context) {
	var urlReq models_http.UrlRequest
	err := DecodeBody(c, &urlReq)
	if ginadap.HandleError(c, err) {
		return
	}

	id, err := GetIdFromContext(c)
	if ginadap.HandleError(c, err) {
		return
	}

	createdUrl, err := h.UrlServices.CreateNewShortUrl(id, urlReq)
	if ginadap.HandleError(c, err) {
		return
	}

	c.JSON(201, gin.H{
		"short_url": createdUrl.ShortURL,
	})
}
