package users

import (
	"github.com/go-openapi/runtime/middleware"

	"otus-arch-hw-3/app/generated/models"
	"otus-arch-hw-3/app/generated/restapi/operations"
)

func (u UserHandler) HandleDeleteUser(params operations.DeleteUserParams) middleware.Responder {
	if err := u.Repository.DeleteUser(params.Body); err != nil {
		errorMessage := err.Error()
		return operations.NewDeleteUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewDeleteUserOK()
}
