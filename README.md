# keycloak-go-vue-auth

## Pre-requisites
- Install Docker
- Install Go


## Run 
- Run `docker compose up`
- Use `http://localhost:8080` to access Keycloak dashboard
- Use `http://localhost:8081` to access Vue app
- Use `http://localhost:8082/api/v1/docs` to access Go app Swagger docs
- To issue tokens, you need client_Id and client_Secret in the realm. Client ID and realm names can be found in `"./backend/config.json"`
- Run to generate API from Swagger. For backend team only.
```bash
make -C ./backend generateAPI 
``` 
