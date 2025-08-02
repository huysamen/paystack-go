package api

import "net/http"

// API is the main client for interacting with the Paystack API.
type API struct {
	Client  *http.Client
	Secret  string
	BaseURL string
}
