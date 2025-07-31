package transfercontrol

import (
	"context"
	"errors"
	"strings"

	"github.com/huysamen/paystack-go/net"
)

// ValidateFinalizeDisableOTPRequest validates the finalize disable OTP request
func ValidateFinalizeDisableOTPRequest(req *FinalizeDisableOTPRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if strings.TrimSpace(req.OTP) == "" {
		return errors.New("otp is required")
	}

	return nil
}

// FinalizeDisableOTP finalizes the request to disable OTP on your transfers
func (c *Client) FinalizeDisableOTP(ctx context.Context, req *FinalizeDisableOTPRequest) (*FinalizeDisableOTPResponse, error) {
	if err := ValidateFinalizeDisableOTPRequest(req); err != nil {
		return nil, err
	}

	resp, err := net.Post[FinalizeDisableOTPRequest, FinalizeDisableOTPResponse](
		ctx, c.client, c.secret, "/transfer/disable_otp_finalize", req, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
