package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"test_iam/generated/swagger/restapi/operations/system"
)

func Health() system.GetHealthHandlerFunc {
	return func(params system.GetHealthParams) middleware.Responder {
		return system.NewGetHealthOK()
	}
}
