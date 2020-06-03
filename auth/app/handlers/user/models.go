package user

import (
	"otus-auth/app/entities"
	"otus-auth/app/generated/models"
)

func buildUser(params *models.UserProperties) entities.User {
	return entities.User{
		Email:        *params.Email,
		PasswordHash: *params.PasswordHash,
		FirstName:    *params.FirstName,
		LastName:     *params.LastName,
	}
}
