package user

import (
	"net/http"
	"otus-auth/app/constants"
	"otus-auth/app/generated/models"
	"otus-auth/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (u UserHandler) HandleSignIn(params operations.SignInParams) middleware.Responder {
	targetUser, err := u.Repository.GetUser(*params.Body.Email)
	if err != nil {
		errorMsg := err.Error()
		return operations.NewSignInBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	if targetUser.PasswordHash != *params.Body.PasswordHash {
		errorMsg := "invalid password"
		return operations.NewSignInInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	session := stringWithCharset(constants.SessionLength)
	u.Sessions[session] = targetUser

	cookie := &http.Cookie{
		Name:     constants.UserSessionKey,
		Value:    session,
		HttpOnly: true,
	}

	return NewSessionCookieResponder(operations.NewSignInOK(), cookie)
}
