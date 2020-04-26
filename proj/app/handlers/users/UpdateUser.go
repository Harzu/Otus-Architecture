package users

import (
	"otus-arch-hw/app/generated/models"
	"otus-arch-hw/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"time"
	"net/http"
)

func (u UserHandler) HandleUpdateUser(params operations.UpdateUserParams) middleware.Responder {
	startedAt, statusCode := time.Now(), http.StatusOK
	defer func() {
		u.afterRequest(params.HTTPRequest.RequestURI, params.HTTPRequest.Method, statusCode, startedAt)
	}()

	user := buildUser(params.Body)
	if err := u.Repository.UpdateUser(user); err != nil {
		errorMessage := err.Error()
		statusCode = http.StatusInternalServerError
		return operations.NewUpdateUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewCreateUserCreated()
}
