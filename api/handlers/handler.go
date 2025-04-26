package handlers

import "github.com/wrtgvr/urlshrt/internal/services"

type Handler struct {
	UserServices *services.UserServices
	UrlServices  *services.UrlServices
	AuthServices *services.AuthServices
}

func NewHandler(userSer *services.UserServices, urlSer *services.UrlServices, authSer *services.AuthServices) *Handler {
	return &Handler{
		UserServices: userSer,
		UrlServices:  urlSer,
	}
}
