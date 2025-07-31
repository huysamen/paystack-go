package transfercontrol

import (
	"net/http"
)

const transferControlBasePath = "/balance"

// Client handles transfer control operations
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new transfer control client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
