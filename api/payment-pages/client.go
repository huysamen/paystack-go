package paymentpages

import (
	"errors"
	"net/http"
)

const paymentPagesBasePath = "/page"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder cannot be nil")

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
