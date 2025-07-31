package paymentrequests

import (
	"errors"
	"strings"
	"time"
)

// ValidateCreatePaymentRequestRequest validates the create payment request
func ValidateCreatePaymentRequestRequest(req *CreatePaymentRequestRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.Customer) == "" {
		return errors.New("customer is required")
	}

	// Either amount or line items must be provided
	if req.Amount == nil && len(req.LineItems) == 0 {
		return errors.New("either amount or line_items must be provided")
	}

	if req.Amount != nil && *req.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Validate line items
	for i, item := range req.LineItems {
		if strings.TrimSpace(item.Name) == "" {
			return errors.New("line item name is required at index " + string(rune(i)))
		}
		if item.Amount <= 0 {
			return errors.New("line item amount must be positive at index " + string(rune(i)))
		}
		if item.Quantity < 0 {
			return errors.New("line item quantity cannot be negative at index " + string(rune(i)))
		}
	}

	// Validate tax items
	for i, tax := range req.Tax {
		if strings.TrimSpace(tax.Name) == "" {
			return errors.New("tax name is required at index " + string(rune(i)))
		}
		if tax.Amount < 0 {
			return errors.New("tax amount cannot be negative at index " + string(rune(i)))
		}
	}

	// Validate due date format if provided
	if req.DueDate != "" {
		if _, err := time.Parse("2006-01-02", req.DueDate); err != nil {
			// Try ISO 8601 format
			if _, err := time.Parse(time.RFC3339, req.DueDate); err != nil {
				return errors.New("due_date must be in YYYY-MM-DD or ISO 8601 format")
			}
		}
	}

	return nil
}

// ValidateUpdatePaymentRequestRequest validates the update payment request
func ValidateUpdatePaymentRequestRequest(req *UpdatePaymentRequestRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.Customer != "" && strings.TrimSpace(req.Customer) == "" {
		return errors.New("customer cannot be empty if provided")
	}

	if req.Amount != nil && *req.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Validate line items
	for i, item := range req.LineItems {
		if strings.TrimSpace(item.Name) == "" {
			return errors.New("line item name is required at index " + string(rune(i)))
		}
		if item.Amount <= 0 {
			return errors.New("line item amount must be positive at index " + string(rune(i)))
		}
		if item.Quantity < 0 {
			return errors.New("line item quantity cannot be negative at index " + string(rune(i)))
		}
	}

	// Validate tax items
	for i, tax := range req.Tax {
		if strings.TrimSpace(tax.Name) == "" {
			return errors.New("tax name is required at index " + string(rune(i)))
		}
		if tax.Amount < 0 {
			return errors.New("tax amount cannot be negative at index " + string(rune(i)))
		}
	}

	// Validate due date format if provided
	if req.DueDate != "" {
		if _, err := time.Parse("2006-01-02", req.DueDate); err != nil {
			// Try ISO 8601 format
			if _, err := time.Parse(time.RFC3339, req.DueDate); err != nil {
				return errors.New("due_date must be in YYYY-MM-DD or ISO 8601 format")
			}
		}
	}

	return nil
}

// ValidateIDOrCode validates an ID or code parameter
func ValidateIDOrCode(idOrCode string) error {
	if strings.TrimSpace(idOrCode) == "" {
		return errors.New("ID or code is required")
	}
	return nil
}

// ValidateCode validates a code parameter
func ValidateCode(code string) error {
	if strings.TrimSpace(code) == "" {
		return errors.New("code is required")
	}
	return nil
}
