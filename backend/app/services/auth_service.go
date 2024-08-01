package services

import (
	"errors"
	"os"

	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/models"
)

type AuthService struct {
	raptor.Service

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

	as.OnInit(func() {
		if as.Key == "" || as.Username == "" || as.Password == "" {
			as.Log.Error("Error creating Auth service; missing auth username, password or key")
			os.Exit(1)
		}
	})

	return as
}

func (as *AuthService) Login(user models.User) (models.User, error) {
	if user.Username == as.Username && user.Password == as.Password {
		user.Key = as.Key
		return user, nil
	}

	return models.User{}, errors.New("invalid username or password")
}
