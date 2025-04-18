package handlers

import "github.com/wrtgvr/urlshrt/internal/services"

type Handler struct {
	UserServices *services.UserServices
	UrlServices  *services.UrlServices
}

func NewHandler(userSer *services.UserServices, urlSer *services.UrlServices) *Handler {
	return &Handler{
		UserServices: userSer,
		UrlServices:  urlSer,
	}
}
