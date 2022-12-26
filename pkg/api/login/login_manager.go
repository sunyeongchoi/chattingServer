package login

import (
	"chatting_back/pkg/login"
	"chatting_back/pkg/models"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIManager struct {
	keyCloak  *login.Keycloak
}

func GetLoginAPIManager(keyCloak *login.Keycloak) *APIManager {
	return &APIManager{
		keyCloak: keyCloak,
	}
}

func (am *APIManager) Login(c echo.Context) error {
	request := models.Login{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	jwt, err := am.keyCloak.Gocloak.Login(context.Background(),
		am.keyCloak.ClientId, am.keyCloak.ClientSecret, am.keyCloak.Realm,
	request.UserName, request.UserPassword)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	rs := &models.LoginResponse{
		AccessToken: jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn: jwt.ExpiresIn,
	}

	return c.JSON(http.StatusOK, rs)
}