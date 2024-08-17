package internal

import openai "github.com/sashabaranov/go-openai"

type OpenAI struct {
	client *openai.Client
}

func NewOpenAI(apiKey string) *OpenAI {
	return &OpenAI{
		client: openai.NewClient(apiKey),
	}
}
