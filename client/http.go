package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type HTTP struct {
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

func NewHTTP(token, baseURL string) *HTTP {
	return &HTTP{token: token, baseURL: baseURL, client: http.Client{}}
}

func (c *HTTP) Do(ctx context.Context, spec requestSpec, output interface{}) (*http.Response, error) {
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

	if output != nil {
		err = json.NewDecoder(httpResp.Body).Decode(&output)
	}

	return httpResp, err
}
