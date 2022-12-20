package login

import (
	"chatting_back/pkg/login"
	"chatting_back/pkg/models"
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

	return c.JSON(http.StatusOK, nil)
}