package repositories

import (
	"otus-arch-hw-3/app/repositories/user"
	"otus-arch-hw-3/app/system/database/psql"
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
