package applepay

import (
	"errors"
	"net/http"
)

const applePayBasePath = "/apple-pay"

var (
	// ErrBuilderRequired is returned when a required builder is nil
	ErrBuilderRequired = errors.New("builder cannot be nil")
)

// Client is the Apple Pay API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Apple Pay API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
