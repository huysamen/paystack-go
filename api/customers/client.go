package customers

import (
	"errors"
	"net/http"
)

const customerBasePath = "/customer"

var (
	// ErrBuilderRequired is returned when a required builder is nil
	ErrBuilderRequired = errors.New("builder cannot be nil")
)

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

func NewClient(client *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  client,
		secret:  secret,
		baseURL: baseURL,
	}
}
