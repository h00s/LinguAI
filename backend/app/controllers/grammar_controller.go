package controllers

import "github.com/go-raptor/raptor/v2"

type GrammarController struct {
	raptor.Controller
}

func (gc *GrammarController) Check(c *raptor.Context) error {
	return c.JSON(raptor.Map{"message": "Check grammar"})
}
