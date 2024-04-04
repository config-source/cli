package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/config-source/cdb"
)

var baseConfigKeyURL = "/api/v1/config-keys"

func (c *Client) GetConfigKey(ctx context.Context, id int) (cdb.ConfigKey, error) {
	var data cdb.ConfigKey

	_, err := c.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/by-id/%d", baseConfigKeyURL, id),
	}, &data)

	return data, err
}

func (c *Client) GetConfigKeyByName(ctx context.Context, name string) (cdb.ConfigKey, error) {
	var data cdb.ConfigKey

	_, err := c.Do(ctx, requestSpec{
		method: "GET",
		url:    fmt.Sprintf("%s/by-name/%s", baseConfigKeyURL, name),
	}, &data)

	return data, err
}

func (c *Client) GetConfigKeyByNameOrID(nameOrID string) (cdb.ConfigKey, error) {
	id, err := strconv.Atoi(nameOrID)
	if err == nil {
		return c.GetConfigKey(context.Background(), id)
	}

	return c.GetConfigKeyByName(context.Background(), nameOrID)
}
