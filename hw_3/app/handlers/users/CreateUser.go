package users

import (
	"github.com/go-openapi/runtime/middleware"

	"otus-arch-hw-3/app/generated/models"
	"otus-arch-hw-3/app/generated/restapi/operations"
)

func (u UserHandler) HandleCreateUser(params operations.CreateUserParams) middleware.Responder {
	user := buildUser(params.Body)
	if _, err := u.Repository.InsertUser(user); err != nil {
		errorMessage := err.Error()
		return operations.NewCreateUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewCreateUserCreated()
}
