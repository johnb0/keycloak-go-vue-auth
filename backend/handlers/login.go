package handlers

import (
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/runtime/middleware"

	"test_iam/generated/swagger/models"
	"test_iam/generated/swagger/restapi/operations/auth"
)

type Auth interface {
	Login() auth.LoginHandlerFunc
	Refresh() auth.RefreshHandlerFunc
}

type KeycloakAuth struct {
	client                            *gocloak.GoCloak
	clientID, clientSecret, realmName string
}

func NewKeycloakAuth(client *gocloak.GoCloak, clientId, clientSecret, realmName string) Auth {
	return &KeycloakAuth{
		client:       client,
		clientID:     clientId,
		clientSecret: clientSecret,
		realmName:    realmName,
	}
}

func (a *KeycloakAuth) Login() auth.LoginHandlerFunc {
	return func(params auth.LoginParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		token, err := a.client.Login(ctx,
			a.clientID,
			a.clientSecret,
			a.realmName,
			params.Login.Username,
			params.Login.Password)
		if err != nil {
			return auth.NewLoginDefault(http.StatusUnauthorized).WithPayload(&models.Error{
				Message: err.Error(),
			})
		}
		return auth.NewLoginOK().
			WithPayload(&models.TokenPair{
				AccessToken:  models.AccessToken(token.AccessToken),
				RefreshToken: models.RefreshToken(token.RefreshToken),
			})
	}
}

func (a *KeycloakAuth) Refresh() auth.RefreshHandlerFunc {
	return func(params auth.RefreshParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		token, err := a.client.RefreshToken(ctx,
			string(params.Refresh),
			a.clientID,
			a.clientSecret,
			a.realmName)
		if err != nil {
			return auth.NewRefreshDefault(http.StatusUnauthorized).WithPayload(&models.Error{
				Message: err.Error(),
			})
		}
		return auth.NewRefreshOK().
			WithPayload(&models.TokenPair{
				AccessToken:  models.AccessToken(token.AccessToken),
				RefreshToken: models.RefreshToken(token.RefreshToken),
			})
	}
}
