package realm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v13"

	"test_iam/config"
)

const masterRealm = "master"
const refClient = "%s-realm"

// TODO: refactor this

func GetRealmClientCred(ctx context.Context, client *gocloak.GoCloak, conf config.Keycloak) (string, string, error) {
	adminName := conf.AdminUser
	adminPassword := conf.AdminPassword
	realm := conf.Realm
	clientName := conf.ClientName

	adminToken, err := client.LoginAdmin(ctx, adminName, adminPassword, masterRealm)
	if err != nil {
		return "", "", fmt.Errorf("failed to login admin: %w", err)
	}

	_, err = client.GetRealm(ctx, adminToken.AccessToken, realm)
	if err != nil {
		apiErr, ok := err.(*gocloak.APIError)
		if !ok {
			return "", "", fmt.Errorf("unknown err: %w", err)
		}
		if apiErr.Code == http.StatusNotFound {
			return createRealmCred(ctx, client, adminToken.AccessToken, conf)
		}
		return "", "", fmt.Errorf("failed to get realm: %w", err)
	}
	keycloakClients, err := client.GetClients(ctx, adminToken.AccessToken, realm, gocloak.GetClientsParams{
		ClientID: &clientName,
	})
	if err != nil {
		return "", "", err
	}
	// todo: fix it
	id := keycloakClients[0].ID
	c, err := client.GetClient(ctx, adminToken.AccessToken, realm, *id)
	if err != nil {
		return "", "", err
	}
	return *c.ClientID,
		*c.Secret,
		nil
}

func createRealmCred(ctx context.Context, client *gocloak.GoCloak, token string, conf config.Keycloak) (string, string, error) {
	realmName := conf.Realm
	clientName := conf.ClientName
	enable := true
	newRealm := gocloak.RealmRepresentation{
		Realm:               &realmName,
		RegistrationAllowed: &enable, // TODO: make it configurable
		Enabled:             &enable,
	}
	_, err := client.CreateRealm(ctx, token, newRealm)
	if err != nil {
		return "", "", fmt.Errorf("failed to create realm: %w", err)
	}

	err = CreateAdminInRealm(ctx, client, token, realmName, conf.AdminUser, conf.AdminPassword)
	if err != nil {
		return "", "", err
	}

	url := fmt.Sprintf("http://0.0.0.0:8080/realms/%s/account/", realmName)
	ns := gocloak.Client{
		ClientID:                  &clientName,
		Name:                      &clientName,
		BaseURL:                   &url,
		DirectAccessGrantsEnabled: &enable,
		ServiceAccountsEnabled:    &enable,
	}
	id, err := client.CreateClient(ctx, token, realmName, ns)
	if err != nil {
		return "", "", fmt.Errorf("failed to create realm client: %w", err)
	}

	c, err := client.GetClient(ctx, token, realmName, id)
	if err != nil {
		return "", "", err
	}
	return *c.Name, *c.Secret, nil
}

func CreateAdminInRealm(ctx context.Context, client *gocloak.GoCloak, token, realm, username, password string) error {
	enable := true
	user := gocloak.User{
		Username: &username,
		Enabled:  &enable,
	}
	id, err := client.CreateUser(ctx, token, realm, user)
	if err != nil {
		return err
	}
	err = client.SetPassword(ctx, token, id, realm, password, false)
	if err != nil {
		return err
	}
	return nil
}
