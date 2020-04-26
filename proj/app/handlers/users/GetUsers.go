package users

import (
	"otus-arch-hw/app/generated/models"
	"otus-arch-hw/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"time"
	"net/http"
)

func (u UserHandler) HandleGetUsers(params operations.UserListParams) middleware.Responder {
	startedAt, statusCode := time.Now(), http.StatusOK
	defer func() {
		u.afterRequest(params.HTTPRequest.RequestURI, params.HTTPRequest.Method, statusCode, startedAt)
	}()

	users, err := u.Repository.GetUsers()
	if err != nil {
		errorMessage := err.Error()
		statusCode = http.StatusInternalServerError
		return operations.NewUserListInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewUserListOK().WithPayload(buildUsersModel(users))
}
