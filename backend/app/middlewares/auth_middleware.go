package middlewares

import (
	"strings"

	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/services"
)

type AuthMiddleware struct {
	raptor.Middleware

	Auth *services.AuthService
}

func (am *AuthMiddleware) New(c *raptor.Context) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.JSONError(raptor.NewErrorUnauthorized("Missing auth header"))
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.JSONError(raptor.NewErrorUnauthorized("Invalid auth header"))
	}

	if authKey := parts[1]; authKey != am.Auth.Key {
		return c.JSONError(raptor.NewErrorUnauthorized("Invalid auth key"))
	}

	return c.Next()
}
