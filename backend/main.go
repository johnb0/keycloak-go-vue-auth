package main

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/loads"

	"test_iam/config"
	"test_iam/generated/swagger/restapi"
	"test_iam/generated/swagger/restapi/operations"
	"test_iam/handlers"
	"test_iam/middleware"
	"test_iam/realm"
)

func main() {
	// config
	ctx := context.Background()
	conf, err := config.GetAppConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalf("Error loading swagger spec: %s", err)
	}
	api := operations.NewTestIamAPI(swaggerSpec)
	api.UseSwaggerUI()

	// keycloak
	client := gocloak.NewClient(conf.Keycloak.Url)
	id, secret, err := realm.GetRealmClientCred(ctx, client, conf.Keycloak)
	if err != nil {
		log.Fatalf("Error getting realm client credentials: %s", err)
	}

	m := middleware.NewKeycloakAuth(client, id, secret, conf.Keycloak)

	// handlers
	api.UserGetUserRolesHandler = handlers.UserRoles()
	api.BearerAuth = m.ParseToken
	api.APIAuthorizer = m

	auth := handlers.NewKeycloakAuth(client, id, secret, conf.Keycloak.Realm)
	api.AuthLoginHandler = auth.Login()
	api.AuthRefreshHandler = auth.Refresh()

	api.Init()
	server := restapi.NewServer(api)

	server.EnabledListeners = []string{"http"}
	server.Host = conf.Server.Host
	server.Port = conf.Server.Port

	server.SetHandler(
		api.Serve(nil),
	)
	// Swagger servers handles signals and gracefully shuts down by itself
	if err = server.Serve(); err != nil {
		log.Fatalf("Error serving: %s", err)
	}

	if errShutdown := server.Shutdown(); errShutdown != nil {
		log.Fatalf("Error shutting down: %s", errShutdown)
	}
}
