package app

import (
	"otus-arch-hw/app/config"
	"otus-arch-hw/app/constants"
	"otus-arch-hw/app/generated/restapi/operations"
	"otus-arch-hw/app/handlers/users"
	"otus-arch-hw/app/repositories"
	"otus-arch-hw/app/services"
	"otus-arch-hw/app/system/database/psql"
	"otus-arch-hw/app/system/healthcheck"
)

type Service struct {
	users.UserHandler
	HealthCheck    healthcheck.HealthCheck
	Services       *services.Services
	PostgresClient psql.Repository
	Repositories   *repositories.Repositories
}

func NewService() Service {
	return Service{}
}

func (sc *Service) ConfigureService() error {
	cfg, err := config.InitConfig(constants.ServiceName)
	if err != nil {
		return err
	}

	postgresClient, err := psql.New(cfg.DBSN)
	if err != nil {
		return err
	}
	sc.PostgresClient = postgresClient
	sc.Repositories = repositories.New(postgresClient)

	sc.Services = services.New()
	sc.HealthCheck = healthcheck.New(postgresClient)

	sc.UserHandler = users.NewHandler(sc.Repositories.GetUserRepo(), sc.Services.Metrics)

	return nil
}

func (sc *Service) SetupSwaggerHandlers(interfaceAPI interface{}) {
	api := interfaceAPI.(*operations.OtusArchHomeworkAPI)
	api.UserListHandler = operations.UserListHandlerFunc(sc.HandleGetUsers)
	api.CreateUserHandler = operations.CreateUserHandlerFunc(sc.HandleCreateUser)
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(sc.HandleUpdateUser)
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(sc.HandleDeleteUser)
}
