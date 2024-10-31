package plans

import (
	"net/http"
)

const (
	planBasePath = "/plan"
)

type Client struct {
	client *http.Client
	secret string
}

func NewClient(secret string, client *http.Client) *Client {
	return &Client{
		secret: secret,
		client: client,
	}
}
