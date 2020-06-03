// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/gorilla/mux"

	"otus-auth/app"
	"otus-auth/app/constants"
	"otus-auth/app/generated/restapi/operations"
)

var srv app.Service

//go:generate swagger generate server --target ../../generated --name OtusArchHomework --spec ../../../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.OtusArchHomeworkAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OtusArchHomeworkAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	srv = app.NewService()
	if err := srv.ConfigureService(); err != nil {
		log.Fatalf("failed to create new services: %+v", err)
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.ServerShutdown = func() {}
	srv.SetupSwaggerHandlers(api)

	return setupCustomRoutes(setupGlobalMiddleware(api.Serve(setupMiddlewares)))
}

func setupCustomRoutes(next http.Handler) http.Handler {
	router := mux.NewRouter()

	router.Handle(constants.HealthCheckURL, http.HandlerFunc(srv.HealthCheck.Check))
	router.NotFoundHandler = next

	return router
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
