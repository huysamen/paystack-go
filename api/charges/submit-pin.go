package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SubmitPIN submits PIN to continue a charge
func (c *Client) SubmitPIN(ctx context.Context, req *SubmitPINRequest) (*SubmitPINResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("submit PIN request cannot be nil")
	}

	if err := validateSubmitPINRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath + "/submit_pin"
	return net.Post[SubmitPINRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateSubmitPINRequest validates the submit PIN request
func validateSubmitPINRequest(req *SubmitPINRequest) error {
	if req.PIN == "" {
		return fmt.Errorf("PIN is required")
	}

	if req.Reference == "" {
		return fmt.Errorf("reference is required")
	}

	// Validate PIN format (should be 4 digits)
	if len(req.PIN) != 4 {
		return fmt.Errorf("PIN must be exactly 4 digits")
	}

	// Check if PIN contains only digits
	for _, char := range req.PIN {
		if char < '0' || char > '9' {
			return fmt.Errorf("PIN must contain only digits")
		}
	}

	return nil
}
