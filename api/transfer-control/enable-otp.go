package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// EnableOTP helps turn OTP requirement back on in the event that a customer wants to stop being able to complete transfers programmatically
// No arguments required.
func (c *Client) EnableOTP(ctx context.Context) (*types.Response[any], error) {
	return net.Post[any, any](
		ctx, c.client, c.secret, "/transfer/enable_otp", nil, c.baseURL,
	)
}
