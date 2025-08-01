package refunds

import (
	"net/http"
)

const refundsBasePath = "/refund"

// Client represents the refunds API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new refunds API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
