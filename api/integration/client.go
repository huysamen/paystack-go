package integration

import (
	"errors"
	"net/http"
)

const integrationBasePath = "/integration"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder cannot be nil")

// Client is the Integration API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Integration API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
