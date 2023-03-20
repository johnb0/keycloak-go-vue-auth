package middleware

import (
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/errors"
)

func (a *KeycloakAuth) Authorize(r *http.Request, auth interface{}) error {
	roles, ok := auth.([]gocloak.Role) // TODO: add some logic here
	if !ok {
		return errors.New(http.StatusForbidden, "invalid token")
	}
	if len(roles) <= 1 {
		return errors.New(http.StatusForbidden, "user has no roles")
	}
	return nil
}
