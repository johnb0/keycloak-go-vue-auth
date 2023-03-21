package handlers

import (
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-openapi/runtime/middleware"

	"test_iam/generated/swagger/restapi/operations/user"
)

func UserRoles() user.GetUserRolesHandlerFunc {
	return func(params user.GetUserRolesParams, auth interface{}) middleware.Responder {
		roles, ok := auth.([]gocloak.Role)
		if !ok {
			return user.NewGetUserRolesDefault(http.StatusInternalServerError).WithStatusCode(http.StatusInternalServerError)
		}
		var response []string
		for _, role := range roles {
			response = append(response, *role.Name)
		}
		return user.NewGetUserRolesOK().WithPayload(response)
	}
}
