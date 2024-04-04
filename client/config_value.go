package client

import (
	"context"
	"fmt"

	"github.com/config-source/cdb"
)

func (ec *Client) GetConfigurationValue(ctx context.Context, environmentName, key string) (cdb.ConfigValue, error) {
	var cv cdb.ConfigValue
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
