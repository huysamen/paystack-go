package api

import "net/http"

type API struct {
	Client  *http.Client
	Secret  string
	BaseURL string
}
