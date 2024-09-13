package services

import (
	"github.com/go-raptor/raptor/v3"
)

type TranslatorService struct {
	raptor.Service

	LLM *LLMService
}
