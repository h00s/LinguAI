package controllers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/models"
	"github.com/h00s/linguai/app/services"
)

type AuthController struct {
	raptor.Controller

	Auth *services.AuthService
}

func (ac *AuthController) Login(c *raptor.Context) error {
	var user models.User
	if err := c.Bind().JSON(&user); err != nil {
		return c.JSONError(err, 400)
	}

	var err error
	if user, err = ac.Auth.Login(user); err != nil {
		return c.JSONError(err, 401)
	}

	return c.JSON(user)
}
