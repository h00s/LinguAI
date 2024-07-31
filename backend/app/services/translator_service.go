package services

import (
	"github.com/go-raptor/raptor/v2"
)

type TranslatorService struct {
	raptor.Service
}

func NewTranslatorService(c *raptor.Config) *TranslatorService {
	ts := &TranslatorService{}

	ts.OnInit(func() {
		// var err error
		// if gs.claude, err = internal.NewClaude(gs.Config.AppConfig["anthropic_key"].(string)); err != nil {
		//	gs.Log.Error("Error creating Claude", "error", err.Error())
		//}
	})

	return ts
}
