package main

import (
	"log"

	"github.com/go-openapi/loads"

	"otus-arch-hw-3/app/config"
	"otus-arch-hw-3/app/constants"
	"otus-arch-hw-3/app/generated/restapi"
	"otus-arch-hw-3/app/generated/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOtusArchHomeworkAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.ConfigureFlags()
	server.ConfigureAPI()

	conf, err := config.InitConfig(constants.ServiceName)
	if err != nil {
		log.Fatalln(err)
	}

	server.Port = conf.HTTPBindPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
