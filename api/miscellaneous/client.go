package miscellaneous

import (
	"errors"
	"net/http"
)

const (
	bankBasePath    = "/bank"
	countryBasePath = "/country"
	statesBasePath  = "/address_verification/states"
)

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder cannot be nil")

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
