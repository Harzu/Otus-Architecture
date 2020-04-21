package app

import (
	"otus-arch-hw-3/app/config"
	"otus-arch-hw-3/app/constants"
	"otus-arch-hw-3/app/generated/restapi/operations"
	"otus-arch-hw-3/app/handlers/users"
	"otus-arch-hw-3/app/repositories"
	"otus-arch-hw-3/app/system/database/psql"
)

type Service struct {
	users.UserHandler
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

	sc.UserHandler = users.NewHandler(sc.Repositories.GetUserRepo())

	return nil
}

func (sc *Service) SetupSwaggerHandlers(interfaceAPI interface{}) {
	api := interfaceAPI.(*operations.OtusArchHomeworkAPI)
	api.UserListHandler = operations.UserListHandlerFunc(sc.HandleGetUsers)
	api.CreateUserHandler = operations.CreateUserHandlerFunc(sc.HandleCreateUser)
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(sc.HandleUpdateUser)
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(sc.HandleDeleteUser)
}
