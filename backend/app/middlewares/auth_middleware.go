package middlewares

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/services"
)

type AuthMiddleware struct {
	raptor.Middleware

	Auth *services.AuthService
}

func (am *AuthMiddleware) New(c *raptor.Context) error {
	if authKey := c.Get("Authorization"); authKey != am.Auth.AuthKey {
		return c.JSONError(raptor.NewErrorUnauthorized("Invalid auth key"))
	}

	return c.Next()
}
