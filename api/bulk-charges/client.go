package bulkcharges

import (
	"net/http"
)

const bulkChargesBasePath = "/bulkcharge"

// Client is the Bulk Charges API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Bulk Charges API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
