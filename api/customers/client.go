package customers

import (
	"net/http"
)

const customerBasePath = "/customer"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(client *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  client,
		secret:  secret,
		baseURL: baseURL,
	}
}
