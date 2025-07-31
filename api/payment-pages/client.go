package paymentpages

import (
	"net/http"
)

const paymentPagesBasePath = "/page"

// Client is the Payment Pages API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Payment Pages API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
