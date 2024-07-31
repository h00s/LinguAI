package services

import (
	"os"

	"github.com/go-raptor/raptor/v2"
)

type AuthService struct {
	raptor.Service

	Key string
}

func NewAuthService(c *raptor.Config) *AuthService {
	as := &AuthService{
		Key: c.AppConfig["auth_key"].(string),
	}

	as.OnInit(func() {
		if as.Key == "" {
			as.Log.Error("Error creating Auth service; missing auth key")
			os.Exit(1)
		}
	})

	return as
}
