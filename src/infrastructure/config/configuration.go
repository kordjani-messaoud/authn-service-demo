package config

import (
	"encoding/json"
	"os"
)

type ConfigParams struct {
	ListenIP   string `json:"listenIP"`
	ListenPort string `json:"listenPort"`
	Keycloak   struct {
		Realm                string `json:"realm"`
		RealmRS256PublickKey string `json:"realmRS256PublicKey"`
		BaseURL              string `json:"baseURL"`
		RestAPI              struct {
			ClientID     string `json:"clientID"`
			ClientSecret string `json:"clientSecret"`
		} `json:"restAPI"`
	} `json:"keycloak"`
}

var GlobalConfigParams ConfigParams

func ExtractConfigParams(path string, configparams *ConfigParams) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(configparams); err != nil {
		return err
	}

	return nil
}

func LoadConfigParamsFromEnv(ConfigParams *ConfigParams) {
	ConfigParams.ListenIP = os.Getenv("LISTEN_IP")
	ConfigParams.ListenPort = os.Getenv("LISTEN_PORT")
	ConfigParams.Keycloak.Realm = os.Getenv("KEYCLOAK_REALM")
	ConfigParams.Keycloak.RealmRS256PublickKey = os.Getenv("KEYCLOAK_REALM_RS256_PUBLIC_KEY")
	ConfigParams.Keycloak.BaseURL = os.Getenv("KEYCLOAK_BASE_URL")
	ConfigParams.Keycloak.RestAPI.ClientID = os.Getenv("KEYCLOAK_CLIENT_ID")
	ConfigParams.Keycloak.RestAPI.ClientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
}
