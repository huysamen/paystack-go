package transaction_splits

import (
	"errors"
	"net/http"
)

const transactionSplitBasePath = "/split"

var (
	// ErrBuilderRequired is returned when a required builder is nil
	ErrBuilderRequired = errors.New("builder cannot be nil")
)

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new transaction splits client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
