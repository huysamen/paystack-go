package virtualterminal

import (
	"errors"
	"net/http"
)

const virtualTerminalBasePath = "/virtual_terminal"

// ErrBuilderRequired is returned when a builder is expected but not provided
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new virtual terminal client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
