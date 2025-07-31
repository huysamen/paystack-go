package verification

import (
	"net/http"
)

const (
	accountResolveBasePath  = "/bank/resolve"
	accountValidateBasePath = "/bank/validate"
	cardBINResolveBasePath  = "/decision/bin"
)

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
