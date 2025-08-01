package directdebit

import (
	"errors"
	"net/http"
)

const directDebitBasePath = "/directdebit"

// ErrBuilderRequired is returned when a required builder is nil
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new direct debit client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
