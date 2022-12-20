package login

import "github.com/Nerzal/gocloak/v9"

type Keycloak struct {
	gocloak 		gocloak.GoCloak // keycloak client
	clientId		string			// clientId specified in keycloak
	clientSecret 	string			// client secret specified in keycloak
	realm 			string			// realm specified in keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		gocloak: gocloak.NewClient("http://localhost:8080"),
		clientId: "test",
		clientSecret: "bvWckUIrzJRIFUo1tZEoZt1vtyczCQgB",
		realm: "test",
	}
}