package handlers

import "github.com/wrtgvr/urlshrt/internal/services"

type Handler struct {
	UserServices  *services.UserServices
	UrlServices   *services.UrlServices
	AuthServices  *services.AuthServices
	TokenServices *services.TokenServices
}

func NewHandler(userS *services.UserServices, urlS *services.UrlServices, authS *services.AuthServices, tokenS *services.TokenServices) *Handler {
	return &Handler{
		userS,
		urlS,
		authS,
		tokenS,
	}
}
