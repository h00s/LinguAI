package services

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/models"
)

type ModelsService struct {
	raptor.Service
}

func (ms *ModelsService) All() []models.Model {
	return []models.Model{
		{
			Code: "gpt-4o",
			Name: "OpenAI GPT-4o",
		},
		{
			Code: "gpt-4o-mini",
			Name: "OpenAI GPT-4o Mini",
		},
		{
			Code: "gpt-35-turbo",
			Name: "OpenAI GPT-3.5 Turbo",
		},
		{
			Code: "claude-35-sonnet",
			Name: "Claude 3.5 Sonnet",
		},
		{
			Code: "claude-3-haiku",
			Name: "Claude 3 Haiku",
		},
	}
}
