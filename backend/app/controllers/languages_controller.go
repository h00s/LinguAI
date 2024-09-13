package controllers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/linguai/app/services"
)

type LanguagesController struct {
	raptor.Controller

	Languages *services.LanguagesService
}

func (lc *LanguagesController) Index(c *raptor.Context) error {
	return c.JSON(lc.Languages.All())
}
