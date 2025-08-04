package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// EnableOTPResponse represents the response from enabling OTP
type EnableOTPResponse = types.Response[any]

// EnableOTP helps turn OTP requirement back on in the event that a customer wants to stop being able to complete transfers programmatically
// No arguments required.
func (c *Client) EnableOTP(ctx context.Context) (*EnableOTPResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, "/transfer/enable_otp", nil, c.BaseURL)
}
