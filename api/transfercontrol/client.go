package transfercontrol

import (
	"errors"
	"net/http"
)

const transferControlBasePath = "/balance"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder is required")

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
