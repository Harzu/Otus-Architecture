package main

import (
	"log"
	"otus-arch-hw/app/config"
	"otus-arch-hw/app/constants"
	"otus-arch-hw/app/generated/restapi"
	"otus-arch-hw/app/generated/restapi/operations"

	"github.com/go-openapi/loads"
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
