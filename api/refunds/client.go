package refunds

import (
	"errors"
	"net/http"
)

const refundsBasePath = "/refund"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder is required - use New*Request() methods to create requests")

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
