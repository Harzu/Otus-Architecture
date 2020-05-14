package user

import (
	"otus-auth/app/generated/models"
	"otus-auth/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (u UserHandler) HandleSignUp(params operations.SignUpParams) middleware.Responder {
	user := buildUser(params.Body)
	if _, err := u.Repository.InsertUser(user); err != nil {
		errorMessage := err.Error()
		return operations.NewSignUpInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	return operations.NewSignUpCreated()
}
