package services

import (
	"errors"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/linguai/app/models"
)

type AuthService struct {
	raptor.Service

	LoginPath string

	Username string
	Password string
	Token    string
}

func NewAuthService(c *raptor.Config) *AuthService {
	as := &AuthService{
		Username: c.AppConfig["auth_username"].(string),
		Password: c.AppConfig["auth_password"].(string),
		Token:    c.AppConfig["auth_token"].(string),
	}
	as.OnInit(as.Init)
	return as
}

func (as *AuthService) Init() error {
	if loginPath, err := as.Routes.Path("AuthController", "Login"); err == nil {
		as.LoginPath = loginPath
	} else {
		return err
	}
	return nil
}

func (as *AuthService) Login(user models.User) (models.User, error) {
	if user.Username == as.Username && user.Password == as.Password {
		user.Token = as.Token
		return user, nil
	}

	return models.User{}, errors.New("invalid username or password")
}
