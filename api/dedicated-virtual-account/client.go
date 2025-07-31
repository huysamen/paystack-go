package dedicatedvirtualaccount

import (
	"fmt"
	"net/http"
)

const dedicatedVirtualAccountBasePath = "/dedicated_account"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new dedicated virtual account client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateCreateRequest(req *CreateDedicatedVirtualAccountRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Customer == "" {
		return fmt.Errorf("customer is required")
	}
	return nil
}

func validateAssignRequest(req *AssignDedicatedVirtualAccountRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if req.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if req.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	if req.PreferredBank == "" {
		return fmt.Errorf("preferred bank is required")
	}
	if req.Country == "" {
		return fmt.Errorf("country is required")
	}
	return nil
}

func validateRequeryRequest(req *RequeryDedicatedAccountRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.AccountNumber == "" {
		return fmt.Errorf("account number is required")
	}
	if req.ProviderSlug == "" {
		return fmt.Errorf("provider slug is required")
	}
	return nil
}

func validateSplitRequest(req *SplitDedicatedAccountTransactionRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Customer == "" {
		return fmt.Errorf("customer is required")
	}
	return nil
}

func validateRemoveSplitRequest(req *RemoveSplitFromDedicatedAccountRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.AccountNumber == "" {
		return fmt.Errorf("account number is required")
	}
	return nil
}

func validateDedicatedAccountID(id string) error {
	if id == "" {
		return fmt.Errorf("dedicated account ID is required")
	}
	return nil
}
