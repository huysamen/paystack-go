package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// SubmitOTP submits OTP to complete a charge
func (c *Client) SubmitOTP(ctx context.Context, req *SubmitOTPRequest) (*SubmitOTPResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("submit OTP request cannot be nil")
	}

	if err := validateSubmitOTPRequest(req); err != nil {
		return nil, err
	}

	url := c.baseURL + chargesBasePath + "/submit_otp"
	return net.Post[SubmitOTPRequest, ChargeData](ctx, c.client, c.secret, url, req)
}

// validateSubmitOTPRequest validates the submit OTP request
func validateSubmitOTPRequest(req *SubmitOTPRequest) error {
	if req.OTP == "" {
		return fmt.Errorf("OTP is required")
	}

	if req.Reference == "" {
		return fmt.Errorf("reference is required")
	}

	// Validate OTP format (typically 4-6 digits)
	if len(req.OTP) < 4 || len(req.OTP) > 6 {
		return fmt.Errorf("OTP must be between 4 and 6 digits")
	}

	// Check if OTP contains only digits
	for _, char := range req.OTP {
		if char < '0' || char > '9' {
			return fmt.Errorf("OTP must contain only digits")
		}
	}

	return nil
}
