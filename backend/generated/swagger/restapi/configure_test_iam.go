// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"test_iam/generated/swagger/restapi/operations"
	"test_iam/generated/swagger/restapi/operations/auth"
	"test_iam/generated/swagger/restapi/operations/system"
	"test_iam/generated/swagger/restapi/operations/user"
)

//go:generate swagger generate server --target ../../../../backend --name TestIam --spec ../../../be.yaml --model-package ./generated/swagger/models --server-package ./generated/swagger/restapi --principal interface{} --exclude-main

func configureFlags(api *operations.TestIamAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TestIamAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.SystemGetHealthHandler == nil {
		api.SystemGetHealthHandler = system.GetHealthHandlerFunc(func(params system.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation system.GetHealth has not yet been implemented")
		})
	}
	if api.UserGetUserRolesHandler == nil {
		api.UserGetUserRolesHandler = user.GetUserRolesHandlerFunc(func(params user.GetUserRolesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserRoles has not yet been implemented")
		})
	}
	if api.AuthLoginHandler == nil {
		api.AuthLoginHandler = auth.LoginHandlerFunc(func(params auth.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Login has not yet been implemented")
		})
	}
	if api.AuthRefreshHandler == nil {
		api.AuthRefreshHandler = auth.RefreshHandlerFunc(func(params auth.RefreshParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.Refresh has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
