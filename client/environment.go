package client

import (
	"context"
	"fmt"
)

type Environment struct {
	Model

	Name         string        `json:"name"`
	Description  string        `json:"description"`
	PromotesToId int64         `json:"promotes_to_id"`
	ConfigValues []ConfigValue `json:"config_values"`
}

func (e Environment) ID() int64 {
	return e.Id
}

func (e Environment) PluralName() string {
	return "environments"
}

type EnvironmentClient struct {
	modelClient[Environment]
}

func NewEnvironmentClient(http *HTTP) EnvironmentClient {
	return EnvironmentClient{
		modelClient: modelClient[Environment]{http},
	}
}

func (ec EnvironmentClient) GetConfigValue(ctx context.Context, environmentName, key string) (ConfigValue, error) {
	var cv ConfigValue
	_, err := ec.HTTP.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("/api/v1/config-values/%s/%s", environmentName, key),
	}, &cv)
	return cv, err
}

func (ec EnvironmentClient) GetByName(ctx context.Context, name string) (Environment, error) {
	var data []Environment

	_, err := ec.modelClient.HTTP.Do(ctx, requestSpec{
		method: "GET",
		url:    ec.modelClient.BaseURL(),
		params: map[string]string{
			"name": name,
		},
	}, &data)

	if len(data) == 0 {
		return Environment{}, fmt.Errorf("no environment with name %s exists", name)
	}

	return data[0], err
}
