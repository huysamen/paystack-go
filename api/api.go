package api

import "net/http"

type API struct {
	Client  *http.Client
	Secret  string
	BaseURL string
	// Optional extra headers to add to each request
	Headers map[string]string
}
