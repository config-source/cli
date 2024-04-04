package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	token   string
	baseURL string
	client  http.Client
}

type requestSpec struct {
	method string
	url    string
	body   interface{}
	params map[string]string
}

func New(token, baseURL string) *Client {
	return &Client{token: token, baseURL: baseURL, client: http.Client{}}
}

func (c *Client) Do(ctx context.Context, spec requestSpec, output interface{}) (*http.Response, error) {
	fullURL := fmt.Sprintf("%s%s", c.baseURL, spec.url)

	req, err := http.NewRequestWithContext(ctx, spec.method, fullURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accepts", "application/json")
	req.Header.Add("Content-Type", "application/json")

	query := url.Values{}
	for key, value := range spec.params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	httpResp, err := c.client.Do(req)
	if err != nil {
		return httpResp, nil
	}
	defer httpResp.Body.Close()

	decoder := json.NewDecoder(httpResp.Body)
	if httpResp.StatusCode >= 400 {
		var errResponse struct {
			Message string `json:"message"`
		}
		err = decoder.Decode(&errResponse)
		if err != nil {
			return nil, err
		}
		err = errors.New(errResponse.Message)
	}

	if output != nil {
		err = decoder.Decode(&output)
	}

	return httpResp, err
}
