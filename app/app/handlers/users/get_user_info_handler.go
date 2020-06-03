package users

import (
	"otus-arch-hw/app/generated/models"
	"otus-arch-hw/app/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func (uh *UserHandler) HandleGetUserInfo(params operations.GetUserInfoParams) middleware.Responder {
	return operations.NewGetUserInfoOK().WithPayload(&models.UserProperties{
		Email:     &params.XUserEmail,
		FirstName: &params.XUserFirstName,
		LastName:  &params.XUserLastName,
	})
}
