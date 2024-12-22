package middlewares

import (
	"strings"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/linguai/app/services"
)

type AuthMiddleware struct {
	raptor.Middleware

	Auth *services.AuthService
}

func (am *AuthMiddleware) New(c *raptor.Context) error {
	if c.Path() == am.Auth.LoginPath {
		return nil
	}
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return raptor.NewErrorUnauthorized("Missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return raptor.NewErrorUnauthorized("Invalid auth header")
	}

	if authKey := parts[1]; authKey != am.Auth.Token {
		return raptor.NewErrorUnauthorized("Invalid auth token")
	}

	return nil
}
