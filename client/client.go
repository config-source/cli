package client

type Client struct {
	http *HTTP
}

func New(token, baseURL string) *Client {
	return &Client{
		http: NewHTTP(token, baseURL),
	}
}

func (c *Client) Environments() EnvironmentClient {
	return NewEnvironmentClient(c.http)
}
