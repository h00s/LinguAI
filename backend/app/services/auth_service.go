package services

import (
	"errors"
	"os"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/linguai/app/models"
)

type AuthService struct {
	raptor.Service

	LoginPath string

	Username string
	Password string
	Key      string
}

func NewAuthService(c *raptor.Config) *AuthService {
	as := &AuthService{
		Username: c.AppConfig["auth_username"].(string),
		Password: c.AppConfig["auth_password"].(string),
		Key:      c.AppConfig["auth_key"].(string),
	}
	as.OnInit(as.Init)
	return as
}

func (as *AuthService) Init() {
	if loginPath, err := as.Routes.Path("AuthController", "Login"); err == nil {
		as.LoginPath = loginPath
	} else {
		as.Log.Error("Error getting login path", "error", err.Error())
		os.Exit(1)
	}
}

func (as *AuthService) Login(user models.User) (models.User, error) {
	if user.Username == as.Username && user.Password == as.Password {
		user.Key = as.Key
		return user, nil
	}

	return models.User{}, errors.New("invalid username or password")
}
