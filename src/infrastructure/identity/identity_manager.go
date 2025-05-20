package identity

import (
	"authn-service-demo/infrastructure/config"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

type IdentityManager struct {
	Realm               string `json:"realm"`
	BaseURL             string `json:"baseURL"`
	RestAPIClientID     string `json:"clientId"`
	RestAPIClientSecret string `json:"clientSecret"`
}

func NewIdentityManager() *IdentityManager {

	return &IdentityManager{
		Realm:               config.GlobalConfigParams.Keycloak.Realm,
		BaseURL:             config.GlobalConfigParams.Keycloak.BaseURL,
		RestAPIClientID:     config.GlobalConfigParams.Keycloak.RestAPI.ClientID,
		RestAPIClientSecret: config.GlobalConfigParams.Keycloak.RestAPI.ClientSecret,
	}
}

func (im *IdentityManager) LoginRestAPIClient(ctx context.Context) (*gocloak.JWT, error) {
	url := "http://localhost:8100"
	client := gocloak.NewClient(url)

	token, err := client.LoginClient(ctx, im.RestAPIClientID, im.RestAPIClientSecret, im.Realm)

	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("error: %v", im))
	}

	return token, err
}

func (im *IdentityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error) {
	token, err := im.LoginRestAPIClient(ctx)
	if err != nil {
		return nil, err
	}
	url := "http://localhost:8100"

	client := gocloak.NewClient(url)

	userID, err := client.CreateUser(ctx, token.AccessToken, im.Realm, user)
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to create the user"))
	}

	err = client.SetPassword(ctx, token.AccessToken, userID, im.Realm, password, false)
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to set password for user"))
	}

	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.Realm, strings.ToLower(role))
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to get realm roles"))
	}

	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.Realm, userID, []gocloak.Role{*roleKeycloak})
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to add role to user"))
	}
	// Add Realm role to user messing

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.Realm, userID)
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to get created user"))
	}

	return userKeycloak, nil
}

func (im *IdentityManager) RetrospectToken(ctx context.Context,
	accessToken string) (*gocloak.IntroSpectTokenResult, error) {

	client := gocloak.NewClient(im.BaseURL)

	rptResult, err := client.RetrospectToken(ctx, accessToken, im.RestAPIClientID, im.RestAPIClientSecret, im.Realm)
	if err != nil {
		return nil, errors.Join(err, errors.New("[identity]: unable to introspect token"))
	}
	return rptResult, nil
}
