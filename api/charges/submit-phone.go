package charges

import (
	"context"
	"fmt"
	"regexp"

	"github.com/huysamen/paystack-go/net"
)

// SubmitPhone submits phone number when requested
func (c *Client) SubmitPhone(ctx context.Context, req *SubmitPhoneRequest) (*SubmitPhoneResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("submit phone request cannot be nil")
	}

	if err := validateSubmitPhoneRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath + "/submit_phone"
	return net.Post[SubmitPhoneRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateSubmitPhoneRequest validates the submit phone request
func validateSubmitPhoneRequest(req *SubmitPhoneRequest) error {
	if req.Phone == "" {
		return fmt.Errorf("phone number is required")
	}

	if req.Reference == "" {
		return fmt.Errorf("reference is required")
	}

	// Validate phone number format (basic validation for Nigerian numbers)
	phoneRegex := regexp.MustCompile(`^(\+234|0)[789][01]\d{8}$`)
	if !phoneRegex.MatchString(req.Phone) {
		return fmt.Errorf("invalid phone number format. Expected Nigerian format like +2348012345678 or 08012345678")
	}

	return nil
}
