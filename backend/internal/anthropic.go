package internal

import (
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic/client/native"
)

type Anthropic struct {
	client *native.Client
}

func NewAnthropic(apiKey string) (*Anthropic, error) {
	if client, err := native.MakeClient(native.Config{APIKey: apiKey}); err == nil {
		return &Anthropic{
			client: client,
		}, nil
	} else {
		return nil, err
	}
}
