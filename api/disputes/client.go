package disputes

import (
	"errors"
	"net/http"
)

const disputesBasePath = "/dispute"

var (
	// ErrBuilderRequired is returned when a required builder is nil
	ErrBuilderRequired = errors.New("builder cannot be nil")
)

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
