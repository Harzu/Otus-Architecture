package user

import (
	"otus-auth/app/constants"
	"otus-auth/app/generated/models"
	"otus-auth/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (u UserHandler) HandleUpdateUser(params operations.UpdateUserParams) middleware.Responder {
	session, err := params.HTTPRequest.Cookie(constants.UserSessionKey)
	if err != nil {
		errorMsg := err.Error()
		return operations.NewUpdateUserBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	tgtUser, ok := u.Sessions[session.Value]
	if !ok {
		errorMsg := "invalid session id"
		return operations.NewUpdateUserBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	newUser := buildUser(params.Body)
	if err = u.Repository.UpdateUser(tgtUser.Email, newUser); err != nil {
		errorMsg := err.Error()
		return operations.NewUpdateUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	u.Sessions[session.Value] = &newUser
	return operations.NewUpdateUserOK()
}
