package users

import (
	userRepo "otus-arch-hw-3/app/repositories/user"
)

type UserHandler struct {
	Repository userRepo.Repository
}

func NewHandler(repo userRepo.Repository) UserHandler {
	return UserHandler{
		Repository: repo,
	}
}
