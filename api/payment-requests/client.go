package paymentrequests

import (
	"net/http"
)

const paymentRequestsBasePath = "/paymentrequest"

// Client is the Payment Requests API client
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
