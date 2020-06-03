package user

import (
	"otus-auth/app/constants"
	"otus-auth/app/generated/models"
	"otus-auth/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (u UserHandler) HandleAuth(params operations.AuthParams) middleware.Responder {
	session, err := params.HTTPRequest.Cookie(constants.UserSessionKey)
	if err != nil {
		errorMsg := err.Error()
		return operations.NewAuthBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	tgtUser, ok := u.Sessions[session.Value]
	if !ok {
		errorMsg := "invalid session id"
		return operations.NewAuthBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	response := operations.NewAuthOK()
	response.SetXUserEmail(tgtUser.Email)
	response.SetXUserFirstName(tgtUser.FirstName)
	response.SetXUserLastName(tgtUser.LastName)
	return response
}
