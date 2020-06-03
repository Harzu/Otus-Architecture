package user

import (
	"otus-auth/app/constants"
	"otus-auth/app/generated/models"
	"otus-auth/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (u UserHandler) HandleLogout(params operations.LogoutParams) middleware.Responder {
	session, err := params.HTTPRequest.Cookie(constants.UserSessionKey)
	if err != nil {
		errorMsg := err.Error()
		return operations.NewLogoutBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	if _, ok := u.Sessions[session.Value]; !ok {
		errorMsg := "invalid session id"
		return operations.NewLogoutBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	delete(u.Sessions, session.Value)
	return operations.NewAuthOK()
}
