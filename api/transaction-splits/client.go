package transaction_splits

import (
	"fmt"
	"net/http"
)

const transactionSplitBasePath = "/split"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new transaction splits client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateCreateRequest(req *TransactionSplitCreateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Type == "" {
		return fmt.Errorf("type is required")
	}
	if req.Type != TransactionSplitTypePercentage && req.Type != TransactionSplitTypeFlat {
		return fmt.Errorf("type must be either 'percentage' or 'flat'")
	}
	if req.Currency.String() == "" {
		return fmt.Errorf("currency is required")
	}
	if len(req.Subaccounts) == 0 {
		return fmt.Errorf("at least one subaccount is required")
	}

	// Validate subaccounts
	for i, subaccount := range req.Subaccounts {
		if subaccount.Subaccount == "" {
			return fmt.Errorf("subaccount[%d]: subaccount code is required", i)
		}
		if subaccount.Share <= 0 {
			return fmt.Errorf("subaccount[%d]: share must be greater than 0", i)
		}
	}

	// Validate bearer type and subaccount combination
	if req.BearerType != nil && *req.BearerType == TransactionSplitBearerTypeSubaccount && req.BearerSubaccount == nil {
		return fmt.Errorf("bearer_subaccount is required when bearer_type is 'subaccount'")
	}

	return nil
}

func validateUpdateRequest(req *TransactionSplitUpdateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	// Validate bearer type and subaccount combination
	if req.BearerType != nil && *req.BearerType == TransactionSplitBearerTypeSubaccount && req.BearerSubaccount == nil {
		return fmt.Errorf("bearer_subaccount is required when bearer_type is 'subaccount'")
	}

	return nil
}

func validateSubaccountAddRequest(req *TransactionSplitSubaccountAddRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Subaccount == "" {
		return fmt.Errorf("subaccount code is required")
	}
	if req.Share <= 0 {
		return fmt.Errorf("share must be greater than 0")
	}

	return nil
}

func validateSubaccountRemoveRequest(req *TransactionSplitSubaccountRemoveRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Subaccount == "" {
		return fmt.Errorf("subaccount code is required")
	}

	return nil
}

func validateTransactionSplitID(id string) error {
	if id == "" {
		return fmt.Errorf("transaction split ID is required")
	}
	return nil
}
