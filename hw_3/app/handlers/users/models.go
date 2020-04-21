package users

import (
	"otus-arch-hw-3/app/entities"
	"otus-arch-hw-3/app/generated/models"
)

func buildUser(params *models.UserProperties) entities.User {
	return entities.User{
		Email:        *params.Email,
		PasswordHash: *params.PasswordHash,
		FirstName:    *params.FirstName,
		LastName:     *params.LastName,
	}
}

func buildUsersModel(users []*entities.User) []*models.UserProperties {
	result := make([]*models.UserProperties, 0, len(users))
	for _, user := range users {
		target := &models.UserProperties{
			Email:        &user.Email,
			FirstName:    &user.FirstName,
			LastName:     &user.LastName,
			PasswordHash: &user.PasswordHash,
		}

		result = append(result, target)
	}

	return result
}
