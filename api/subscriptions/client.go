package subscriptions

import (
	"errors"
	"net/http"
)

const subscriptionBasePath = "/subscription"

// ErrBuilderRequired is returned when a builder is expected but not provided
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
