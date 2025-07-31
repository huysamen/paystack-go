package plans

import (
	"net/http"
)

const (
	planBasePath = "/plan"
)

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	if baseURL == "" {
		baseURL = "https://api.paystack.co"
	}

	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
