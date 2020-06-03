package app

import (
	"otus-arch-hw/app/generated/restapi/operations"
	"otus-arch-hw/app/handlers/users"
	"otus-arch-hw/app/system/healthcheck"
)

type Service struct {
	users.UserHandler
	HealthCheck healthcheck.HealthCheck
}

func NewService() Service {
	return Service{}
}

func (sc *Service) ConfigureService() error {
	sc.HealthCheck = healthcheck.New()
	sc.UserHandler = users.NewHandler()

	return nil
}

func (sc *Service) SetupSwaggerHandlers(interfaceAPI interface{}) {
	api := interfaceAPI.(*operations.OtusArchHomeworkAPI)
	api.GetUserInfoHandler = operations.GetUserInfoHandlerFunc(sc.HandleGetUserInfo)
}
