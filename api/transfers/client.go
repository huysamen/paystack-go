package transfers

import (
	"net/http"
)

const transferBasePath = "/transfer"

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
