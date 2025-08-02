package applepay

import (
	"net/http"

	"github.com/huysamen/paystack-go/api"
)

const (
	basePath       = "/apple-pay"
	listPath       = basePath + "/domain"
	registerPath   = basePath + "/domain"
	unregisterPath = basePath + "/domain"
)

// Client is the Apple Pay API client
type Client api.API

// NewClient creates a new Apple Pay API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		Client:  httpClient,
		Secret:  secret,
		BaseURL: baseURL,
	}
}
