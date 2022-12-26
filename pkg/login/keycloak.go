package login

import "github.com/Nerzal/gocloak/v9"

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in keycloak
	ClientSecret string          // client secret specified in keycloak
	Realm        string          // realm specified in keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient("http://localhost:8080"),
		ClientId:     "test",
		ClientSecret: "bvWckUIrzJRIFUo1tZEoZt1vtyczCQgB",
		Realm:        "test",
	}
}