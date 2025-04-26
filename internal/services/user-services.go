package services

import rep "github.com/wrtgvr/urlshrt/internal/repository"

type UserServices struct {
	UserRepo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{repo}
}
