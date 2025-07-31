package paymentpages

import (
	"errors"
	"strings"
)

// ValidateCreatePaymentPageRequest validates the create payment page request
func ValidateCreatePaymentPageRequest(req *CreatePaymentPageRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if req.Type != "" {
		validTypes := []string{"payment", "subscription", "product", "plan"}
		isValidType := false
		for _, validType := range validTypes {
			if req.Type == validType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			return errors.New("type must be one of: payment, subscription, product, plan")
		}
	}

	if req.Type == "subscription" && req.Plan == "" {
		return errors.New("plan is required when type is subscription")
	}

	if req.FixedAmount != nil && *req.FixedAmount && req.Amount == nil {
		return errors.New("amount is required when fixed_amount is true")
	}

	if req.Amount != nil && *req.Amount < 0 {
		return errors.New("amount must be non-negative")
	}

	return nil
}

// ValidateUpdatePaymentPageRequest validates the update payment page request
func ValidateUpdatePaymentPageRequest(req *UpdatePaymentPageRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.Name != "" && strings.TrimSpace(req.Name) == "" {
		return errors.New("name cannot be empty if provided")
	}

	if req.Amount != nil && *req.Amount < 0 {
		return errors.New("amount must be non-negative")
	}

	return nil
}

// ValidateAddProductsToPageRequest validates the add products to page request
func ValidateAddProductsToPageRequest(req *AddProductsToPageRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if len(req.Product) == 0 {
		return errors.New("at least one product ID is required")
	}

	for _, productID := range req.Product {
		if productID <= 0 {
			return errors.New("all product IDs must be positive integers")
		}
	}

	return nil
}

// ValidateIDOrSlug validates an ID or slug parameter
func ValidateIDOrSlug(idOrSlug string) error {
	if strings.TrimSpace(idOrSlug) == "" {
		return errors.New("ID or slug is required")
	}
	return nil
}

// ValidateSlug validates a slug parameter
func ValidateSlug(slug string) error {
	if strings.TrimSpace(slug) == "" {
		return errors.New("slug is required")
	}
	return nil
}

// ValidatePageID validates a page ID parameter
func ValidatePageID(pageID int) error {
	if pageID <= 0 {
		return errors.New("page ID must be a positive integer")
	}
	return nil
}
