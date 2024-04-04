package utils

import (
	"github.com/config-source/cli/client"
	"github.com/config-source/cli/config"
)

var apiClient *client.Client

func GetClient() *client.Client {
	if apiClient == nil {
		apiClient = client.New(config.Current.Token, config.Current.BaseURL)
	}

	return apiClient
}
