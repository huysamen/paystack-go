package settlements

import (
	"errors"
	"net/http"
)

const settlementBasePath = "/settlement"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder is required - use New*Request() methods to create requests")

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
