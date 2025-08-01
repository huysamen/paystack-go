package paymentrequests

import (
	"errors"
	"net/http"
)

const paymentRequestsBasePath = "/paymentrequest"

// ErrBuilderRequired is returned when a required builder parameter is missing
var ErrBuilderRequired = errors.New("builder is required")

// Client is the API client for the Payment Requests API
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Payment Requests API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
