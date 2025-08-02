package bulkcharges

import (
	"net/http"

	"github.com/huysamen/paystack-go/api"
)

const (
	basePath         = "/bulkcharge"
	pausePath        = basePath + "/pause"
	resumePath       = basePath + "/resume"
	fetchChargesPath = "/charges"
)

// Client is the Bulk Charges API client
type Client api.API

// NewClient creates a new Bulk Charges API client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		Client:  httpClient,
		Secret:  secret,
		BaseURL: baseURL,
	}
}
