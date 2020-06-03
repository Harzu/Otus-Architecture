package user

import (
	"otus-auth/app/entities"
	userRepo "otus-auth/app/repositories/user"
)

type UserHandler struct {
	Repository userRepo.Repository
	Sessions   map[string]*entities.User
}

func NewHandler(repo userRepo.Repository) UserHandler {
	return UserHandler{
		Repository: repo,
		Sessions:   make(map[string]*entities.User),
	}
}
