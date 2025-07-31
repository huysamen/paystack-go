package miscellaneous

import (
	"net/http"
)

const (
	bankBasePath    = "/bank"
	countryBasePath = "/country"
	statesBasePath  = "/address_verification/states"
)

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new miscellaneous client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
