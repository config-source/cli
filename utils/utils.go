package utils

import "github.com/config-source/cli/client"

var apiClient *client.Client

func GetClient() *client.Client {
	if apiClient == nil {
		// TODO: get token and url from config
		apiClient = client.New("", "http://localhost:8080")
	}

	return apiClient
}
