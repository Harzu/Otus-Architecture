package users

import (
	"github.com/go-openapi/runtime/middleware"

	"otus-arch-hw-3/app/generated/models"
	"otus-arch-hw-3/app/generated/restapi/operations"
)

func (u UserHandler) HandleGetUsers(params operations.UserListParams) middleware.Responder {
	users, err := u.Repository.GetUsers()
	if err != nil {
		errorMessage := err.Error()
		return operations.NewUserListInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewUserListOK().WithPayload(buildUsersModel(users))
}
