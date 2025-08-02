package verification

import (
	"errors"
	"net/http"
)

const (
	accountResolveBasePath  = "/bank/resolve"
	accountValidateBasePath = "/bank/validate"
	cardBINResolveBasePath  = "/decision/bin"
)

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new verification client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
