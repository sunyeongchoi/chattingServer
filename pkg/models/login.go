package models

type Login struct {
	UserName		string	`json:"userName"`
	UserPassword	string	`json:"userPassword"`
}

type LoginResponse struct {
	AccessToken		string	`json:"accessToken"`
	RefreshToken	string	`json:"refreshToken"`
	ExpiresIn		int		`json:"expiresIn"`
}
