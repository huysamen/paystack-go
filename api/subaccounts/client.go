package subaccounts

import (
	"fmt"
	"net/http"
)

const subaccountBasePath = "/subaccount"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new subaccounts client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateCreateRequest(req *SubaccountCreateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.BusinessName == "" {
		return fmt.Errorf("business_name is required")
	}
	if req.BankCode == "" {
		return fmt.Errorf("settlement_bank is required")
	}
	if req.AccountNumber == "" {
		return fmt.Errorf("account_number is required")
	}
	if req.PercentageCharge < 0 || req.PercentageCharge > 100 {
		return fmt.Errorf("percentage_charge must be between 0 and 100")
	}
	return nil
}

func validateUpdateRequest(req *SubaccountUpdateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.BusinessName == "" {
		return fmt.Errorf("business_name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("description is required")
	}
	if req.PercentageCharge != nil && (*req.PercentageCharge < 0 || *req.PercentageCharge > 100) {
		return fmt.Errorf("percentage_charge must be between 0 and 100")
	}
	return nil
}
