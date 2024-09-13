package services

import "github.com/go-raptor/raptor/v3"

type GrammarService struct {
	raptor.Service

	LLM *LLMService
}
