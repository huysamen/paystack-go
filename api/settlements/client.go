package settlements

import (
	"net/http"
)

const settlementBasePath = "/settlement"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new settlements client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
