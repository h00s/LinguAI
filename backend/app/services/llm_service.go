package services

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/internal"
)

type LLMService struct {
	raptor.Service

	OpenAI    *internal.OpenAI
	Anthropic *internal.Anthropic
}

func NewLLMService(c *raptor.Config) *LLMService {
	llms := &LLMService{}

	llms.OnInit(func() {
		var err error
		if llms.Anthropic, err = internal.NewAnthropic(llms.Config.AppConfig["anthropic_key"].(string)); err != nil {
			llms.Log.Error("Error creating Claude", "error", err.Error())
		}

		llms.OpenAI = internal.NewOpenAI(llms.Config.AppConfig["openai_key"].(string))
	})

	return llms
}
