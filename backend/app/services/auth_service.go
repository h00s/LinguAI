package services

import (
	"os"

	"github.com/go-raptor/raptor/v2"
)

type AuthService struct {
	raptor.Service

	AuthKey string
}

func NewAuthService(c *raptor.Config) *AuthService {
	as := &AuthService{
		AuthKey: c.AppConfig["auth_key"].(string),
	}

	as.OnInit(func() {
		if as.AuthKey == "" {
			as.Log.Error("Error creating Auth service; missing auth key")
			os.Exit(1)
		}
	})

	return as
}
