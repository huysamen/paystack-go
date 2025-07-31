package directdebit

import (
	"fmt"
	"net/http"
)

const directDebitBasePath = "/directdebit"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new direct debit client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateTriggerActivationChargeRequest(req *TriggerActivationChargeRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if len(req.CustomerIDs) == 0 {
		return fmt.Errorf("customer IDs are required")
	}
	for i, customerID := range req.CustomerIDs {
		if customerID <= 0 {
			return fmt.Errorf("invalid customer ID at index %d: must be positive", i)
		}
	}
	return nil
}
