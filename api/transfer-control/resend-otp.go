package transfercontrol

import (
	"context"
	"errors"
	"strings"

	"github.com/huysamen/paystack-go/net"
)

// ValidateResendOTPRequest validates the resend OTP request
func ValidateResendOTPRequest(req *ResendOTPRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.TransferCode) == "" {
		return errors.New("transfer_code is required")
	}

	if strings.TrimSpace(req.Reason) == "" {
		return errors.New("reason is required")
	}

	// Validate reason values
	validReasons := []string{"resend_otp", "transfer"}
	isValidReason := false
	for _, validReason := range validReasons {
		if req.Reason == validReason {
			isValidReason = true
			break
		}
	}

	if !isValidReason {
		return errors.New("reason must be either 'resend_otp' or 'transfer'")
	}

	return nil
}

// ResendOTP generates a new OTP and sends to customer in the event they are having trouble receiving one
func (c *Client) ResendOTP(ctx context.Context, req *ResendOTPRequest) (*ResendOTPResponse, error) {
	if err := ValidateResendOTPRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[ResendOTPRequest, ResendOTPResponse](
		ctx, c.client, c.secret, "/transfer/resend_otp", req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
