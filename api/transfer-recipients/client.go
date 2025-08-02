package transfer_recipients

import (
	"errors"
	"fmt"
	"net/http"
)

const transferRecipientBasePath = "/transferrecipient"

// ErrBuilderRequired is returned when a builder is required but not provided
var ErrBuilderRequired = errors.New("builder is required")

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new transfer recipients client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateCreateRequest(req *TransferRecipientCreateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Type == "" {
		return fmt.Errorf("type is required")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Type != "authorization" {
		if req.AccountNumber == "" {
			return fmt.Errorf("account_number is required for type %s", req.Type)
		}
		if req.BankCode == "" {
			return fmt.Errorf("bank_code is required for type %s", req.Type)
		}
	}
	return nil
}

func validateBulkCreateRequest(req *BulkCreateTransferRecipientRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if len(req.Batch) == 0 {
		return fmt.Errorf("batch cannot be empty")
	}

	for i, item := range req.Batch {
		if item.Type == "" {
			return fmt.Errorf("batch[%d]: type is required", i)
		}
		if item.Name == "" {
			return fmt.Errorf("batch[%d]: name is required", i)
		}
		if item.Type != "authorization" {
			if item.AccountNumber == "" {
				return fmt.Errorf("batch[%d]: account_number is required for type %s", i, item.Type)
			}
			if item.BankCode == "" {
				return fmt.Errorf("batch[%d]: bank_code is required for type %s", i, item.Type)
			}
		}
	}
	return nil
}

func validateUpdateRequest(req *TransferRecipientUpdateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
