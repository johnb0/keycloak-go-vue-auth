# keycloak-go-vue-auth

## Pre-requisites
- Install Docker

## Run 
- Run `docker compose up`
- Use `http://localhost:8080` to access Keycloak dashboard
- Use `http://localhost:8081` to access Vue app
- Use `http://localhost:8082/api/v1/docs` to access Go app Swagger docs
- To issue tokens, you need client_Id and client_Secret in the realm. Client ID and realm names can be found in `"./backend/config.json"`
