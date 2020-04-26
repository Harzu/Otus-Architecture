package users

import (
	"otus-arch-hw/app/generated/models"
	"otus-arch-hw/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"time"
	"net/http"
)

func (u UserHandler) HandleCreateUser(params operations.CreateUserParams) middleware.Responder {
	startedAt, statusCode := time.Now(), http.StatusOK
	defer func() {
		u.afterRequest(params.HTTPRequest.RequestURI, params.HTTPRequest.Method, statusCode, startedAt)
	}()

	user := buildUser(params.Body)
	if _, err := u.Repository.InsertUser(user); err != nil {
		errorMessage := err.Error()
		statusCode = http.StatusInternalServerError
		return operations.NewCreateUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewCreateUserCreated()
}
