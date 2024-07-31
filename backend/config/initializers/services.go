package initializers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/services"
)

func Services(c *raptor.Config) raptor.Services {
	return raptor.Services{
		services.NewAuthService(c),
		services.NewGrammarService(c),
		services.NewTranslatorService(c),
	}
}
