package controllers

import "github.com/go-raptor/raptor/v3"

type GrammarController struct {
	raptor.Controller
}

func (gc *GrammarController) Check(c *raptor.Context) error {
	return c.JSON(raptor.Map{"message": "Check grammar"})
}
