package services

import (
	"github.com/go-raptor/raptor/v2"
)

type TranslatorService struct {
	raptor.Service

	LLM *LLMService
}
