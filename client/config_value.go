package client

import (
	"context"
	"fmt"

	"github.com/config-source/cdb"
)

func (ec *Client) GetConfigurationValue(ctx context.Context, environmentName, key string) (*cdb.ConfigValue, error) {
	var cv *cdb.ConfigValue
	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("/api/v1/config-values/%s/%s", environmentName, key),
	}, &cv)
	return cv, err
}

func (ec *Client) GetConfiguration(ctx context.Context, environmentName string) ([]cdb.ConfigValue, error) {
	var values []cdb.ConfigValue
	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("/api/v1/config-values/%s", environmentName),
	}, &values)
	return values, err
}

func (ec *Client) SetConfiguration(ctx context.Context, value *cdb.ConfigValue) (*cdb.ConfigValue, error) {
	var setValue *cdb.ConfigValue
	_, err := ec.Do(ctx, requestSpec{
		method: "POST",
		url:    "/api/v1/config-values",
		body:   value,
	}, &setValue)
	return setValue, err
}

func (ec *Client) SetConfigurationValue(ctx context.Context, env string, key string, value *cdb.ConfigValue) (*cdb.ConfigValue, error) {
	var setValue *cdb.ConfigValue
	_, err := ec.Do(ctx, requestSpec{
		method: "POST",
		url:    fmt.Sprintf("/api/v1/config-values/%s/%s", env, key),
		body:   value,
	}, &setValue)
	return setValue, err
}
