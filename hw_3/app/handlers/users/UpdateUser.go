package users

import (
	"github.com/go-openapi/runtime/middleware"

	"otus-arch-hw-3/app/generated/models"
	"otus-arch-hw-3/app/generated/restapi/operations"
)

func (u UserHandler) HandleUpdateUser(params operations.UpdateUserParams) middleware.Responder {
	user := buildUser(params.Body)
	if err := u.Repository.UpdateUser(user); err != nil {
		errorMessage := err.Error()
		return operations.NewUpdateUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewCreateUserCreated()
}
