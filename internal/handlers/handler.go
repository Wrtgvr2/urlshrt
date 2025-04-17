package handlers

import "github.com/wrtgvr/urlshrt/internal/services"

type Handler struct {
	UserServices services.UserServices
	UrlServices  services.UrlServices
}
