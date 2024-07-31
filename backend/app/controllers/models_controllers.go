package controllers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/services"
)

type ModelsController struct {
	raptor.Controller

	Models *services.ModelsService
}

func (mc *ModelsController) All(c *raptor.Context) error {
	return c.JSON(mc.Models.All())
}
