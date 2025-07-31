package charges

import (
	"net/http"
)

const chargesBasePath = "/charge"

// Client is the Charges API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Charges API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
