package plans

import (
	"errors"
	"net/http"
)

const (
	planBasePath = "/plan"
)

// ErrBuilderRequired is returned when a required builder parameter is missing
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	if baseURL == "" {
		baseURL = "https://api.paystack.co"
	}

	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
