package services

import "github.com/go-raptor/raptor/v2"

type GrammarService struct {
	raptor.Service

	LLM *LLMService
}
