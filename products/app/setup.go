package app

import (
	"otus-products/app/config"
	"otus-products/app/constants"
	"otus-products/app/generated/restapi/operations"
	"otus-products/app/handlers/products"
	"otus-products/app/repositories"
	"otus-products/app/system/database/psql"
	"otus-products/app/system/database/redis"
	"otus-products/app/system/healthcheck"
)

type Service struct {
	products.ProductsHandler
	HealthCheck     healthcheck.HealthCheck
	PostgresClient  psql.Repository
	RedisRepository redis.Repository
	Repositories    *repositories.Repositories
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

	redisRepo, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}

	sc.PostgresClient = postgresClient
	sc.RedisRepository = redisRepo
	sc.Repositories = repositories.New(sc.PostgresClient)
	sc.HealthCheck = healthcheck.New(sc.PostgresClient, sc.RedisRepository)
	sc.ProductsHandler = products.NewHandler(sc.Repositories.GetProductsRepo(), sc.RedisRepository)

	return nil
}

func (sc *Service) SetupSwaggerHandlers(interfaceAPI interface{}) {
	api := interfaceAPI.(*operations.OtusArchHomeworkAPI)
	api.GetProductsHandler = operations.GetProductsHandlerFunc(sc.ProductsHandler.GetProducts)
}
