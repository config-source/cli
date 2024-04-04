package client

import (
	"context"
	"fmt"

	"github.com/config-source/cdb"
)

var baseEnvURL = "/api/v1/environments"

func (ec *Client) GetEnvironmentByName(ctx context.Context, name string) (cdb.Environment, error) {
	var data cdb.Environment

	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/by-name/%s", baseEnvURL, name),
	}, &data)

	return data, err
}

func (ec *Client) GetEnvironment(ctx context.Context, id int) (cdb.Environment, error) {
	var data cdb.Environment

	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/by-id/%d", baseEnvURL, id),
	}, &data)

	return data, err
}

func (ec *Client) CreateEnvironment(ctx context.Context, env cdb.Environment) (cdb.Environment, error) {
	var data cdb.Environment

	_, err := ec.Do(ctx, requestSpec{
		method: "POST",
		url:    baseEnvURL,
		body:   env,
	}, &data)

	return data, err
}

func (ec *Client) ListEnvironments(ctx context.Context) ([]cdb.Environment, error) {
	var data []cdb.Environment

	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    baseEnvURL,
	}, &data)

	return data, err
}

func (ec *Client) GetEnvironmentTree(ctx context.Context) ([]cdb.EnvTree, error) {
	var data []cdb.EnvTree

	_, err := ec.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/tree", baseEnvURL),
	}, &data)

	return data, err
}
