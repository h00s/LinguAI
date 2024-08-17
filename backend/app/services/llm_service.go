package services

import "github.com/go-raptor/raptor/v2"

type LLMService struct {
	raptor.Service
}

func NewLLMService(c *raptor.Config) *LLMService {
	llms := &LLMService{}

	llms.OnInit(func() {
		// var err error
		// if gs.claude, err = internal.NewClaude(gs.Config.AppConfig["anthropic_key"].(string)); err != nil {
		//	gs.Log.Error("Error creating Claude", "error", err.Error())
		//}
	})

	return llms
}
