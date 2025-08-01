package disputes

import (
	"net/http"
)

const disputesBasePath = "/dispute"

// ErrBuilderRequired is returned when a builder is required but not provided
const ErrBuilderRequired = "builder is required"

// Client represents the disputes API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new disputes API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
