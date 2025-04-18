package services

import (
	rep "github.com/wrtgvr/urlshrt/internal/repository"
)

type UserServices struct {
	Repo rep.UserRepo
}

func NewUserServices(repo rep.UserRepo) UserServices {
	return UserServices{Repo: repo}
}

func (s *UserServices) Login() {
	// Get user my username
	// if no user - return err
	// if user exists check given password with hash in DB
	// if it's ok then create tokens and return it
	// that's all(?)
	//s.Repo.GetUserByUsername()
}
