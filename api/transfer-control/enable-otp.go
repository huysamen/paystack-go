package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// EnableOTP helps turn OTP requirement back on in the event that a customer wants to stop being able to complete transfers programmatically
// No arguments required.
func (c *Client) EnableOTP(ctx context.Context) (*EnableOTPResponse, error) {
	resp, err := net.Post[any, EnableOTPResponse](
		ctx, c.client, c.secret, "/transfer/enable_otp", nil, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
