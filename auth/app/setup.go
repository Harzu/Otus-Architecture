package app

import (
	"otus-auth/app/config"
	"otus-auth/app/constants"
	"otus-auth/app/generated/restapi/operations"
	"otus-auth/app/handlers/user"
	"otus-auth/app/repositories"
	"otus-auth/app/system/database/psql"
	"otus-auth/app/system/healthcheck"
)

type Service struct {
	user.UserHandler
	HealthCheck    healthcheck.HealthCheck
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
	sc.Repositories = repositories.New(sc.PostgresClient)
	sc.HealthCheck = healthcheck.New(sc.PostgresClient)
	sc.UserHandler = user.NewHandler(sc.Repositories.GetUserRepo())

	return nil
}

func (sc *Service) SetupSwaggerHandlers(interfaceAPI interface{}) {
	api := interfaceAPI.(*operations.OtusArchHomeworkAPI)
	api.AuthHandler = operations.AuthHandlerFunc(sc.HandleAuth)
	api.SignUpHandler = operations.SignUpHandlerFunc(sc.HandleSignUp)
	api.SignInHandler = operations.SignInHandlerFunc(sc.HandleSignIn)
	api.LogoutHandler = operations.LogoutHandlerFunc(sc.HandleLogout)
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(sc.HandleUpdateUser)
}
