package repositories

import (
	"otus-auth/app/repositories/user"
	"otus-auth/app/system/database/psql"
)

type Repositories struct {
	userRepo user.Repository
}

func New(client psql.Repository) *Repositories {
	return &Repositories{
		userRepo: user.NewRepository(client),
	}
}

func (r *Repositories) GetUserRepo() user.Repository {
	return r.userRepo
}
