package services

import (
	rep "github.com/wrtgvr/urlshrt/internal/repository"
)

type UrlServices struct {
	Repo rep.UrlRepo
}

func NewUrlServices(repo rep.UrlRepo) UrlServices {
	return UrlServices{Repo: repo}
}
