package middleware

import (
	"context"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/errors"
	"github.com/golang-jwt/jwt/v4"

	"test_iam/config"
)

const subClaim = "sub"

type KeycloakAuth struct {
	client       *gocloak.GoCloak
	clientID     string
	clientSecret string
	realmName    string
	adminUser    string
	adminPass    string
}

func NewKeycloakAuth(client *gocloak.GoCloak, clientID, clientSecret string, conf config.Keycloak) *KeycloakAuth {
	return &KeycloakAuth{
		client:       client,
		clientID:     clientID,
		clientSecret: clientSecret,
		realmName:    conf.Realm,
		adminUser:    conf.AdminUser,
		adminPass:    conf.AdminPassword,
	}
}

func (a *KeycloakAuth) ParseToken(token string) (interface{}, error) {
	ctx := context.Background()
	i, err := a.client.RetrospectToken(ctx, token, a.clientID, a.clientSecret, a.realmName)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, err.Error()) // todo: remove error from return
	}
	if !*i.Active {
		return nil, errors.New(http.StatusUnauthorized, "invalid token")
	}
	_, claims, err := a.client.DecodeAccessToken(ctx, token, a.realmName)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, err.Error())
	}
	return a.GetUserRoles(claims)
}

func (a *KeycloakAuth) GetUserRoles(claims *jwt.MapClaims) ([]gocloak.Role, error) {
	ctx := context.Background()
	m := map[string]interface{}(*claims)
	userID, ok := m[subClaim].(string)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, "invalid token")
	}
	t, err := a.client.LoginAdmin(ctx, a.adminUser, a.adminPass, a.realmName)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, err.Error())
	}
	roles, err := a.client.GetRealmRolesByUserID(ctx, t.AccessToken, a.realmName, userID)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, err.Error())
	}
	rolesWithAttributes := make([]gocloak.Role, 0)
	for _, role := range roles {
		r, err := a.client.GetRealmRole(ctx, t.AccessToken, a.realmName, *role.Name)
		if err != nil {
			return nil, errors.New(http.StatusUnauthorized, err.Error())
		}
		rolesWithAttributes = append(rolesWithAttributes, *r)
	}
	return rolesWithAttributes, nil
}
