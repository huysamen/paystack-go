package products

import (
	"net/http"
)

const productsBasePath = "/product"

// Client is the Products API client
type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new Products API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}
