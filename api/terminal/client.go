package terminal

import (
	"fmt"
	"net/http"
)

const terminalBasePath = "/terminal"

type Client struct {
	client  *http.Client
	secret  string
	baseURL string
}

// NewClient creates a new terminal client
func NewClient(httpClient *http.Client, secret, baseURL string) *Client {
	return &Client{
		client:  httpClient,
		secret:  secret,
		baseURL: baseURL,
	}
}

// Validation functions

func validateSendEventRequest(req *TerminalSendEventRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.Type == "" {
		return fmt.Errorf("event type is required")
	}
	if req.Type != TerminalEventTypeInvoice && req.Type != TerminalEventTypeTransaction {
		return fmt.Errorf("event type must be either 'invoice' or 'transaction'")
	}
	if req.Action == "" {
		return fmt.Errorf("event action is required")
	}

	// Validate action based on type
	if req.Type == TerminalEventTypeInvoice {
		if req.Action != TerminalEventActionProcess && req.Action != TerminalEventActionView {
			return fmt.Errorf("for invoice type, action must be either 'process' or 'view'")
		}
	} else if req.Type == TerminalEventTypeTransaction {
		if req.Action != TerminalEventActionProcess && req.Action != TerminalEventActionPrint {
			return fmt.Errorf("for transaction type, action must be either 'process' or 'print'")
		}
	}

	if req.Data == nil {
		return fmt.Errorf("event data is required")
	}

	return nil
}

func validateUpdateRequest(req *TerminalUpdateRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	// At least one field should be provided for update
	if req.Name == nil && req.Address == nil {
		return fmt.Errorf("at least one field (name or address) must be provided for update")
	}

	return nil
}

func validateCommissionRequest(req *TerminalCommissionRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.SerialNumber == "" {
		return fmt.Errorf("serial number is required")
	}

	return nil
}

func validateDecommissionRequest(req *TerminalDecommissionRequest) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}
	if req.SerialNumber == "" {
		return fmt.Errorf("serial number is required")
	}

	return nil
}

func validateTerminalID(terminalID string) error {
	if terminalID == "" {
		return fmt.Errorf("terminal ID is required")
	}
	return nil
}

func validateEventID(eventID string) error {
	if eventID == "" {
		return fmt.Errorf("event ID is required")
	}
	return nil
}
