package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// DisableOTP is used in the event that you want to be able to complete transfers programmatically without use of OTPs
// No arguments required. You will get an OTP to complete the request.
func (c *Client) DisableOTP(ctx context.Context) (*DisableOTPResponse, error) {
	resp, err := net.Post[any, DisableOTPResponse](
		ctx, c.client, c.secret, "/transfer/disable_otp", nil, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
