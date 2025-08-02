package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// DisableOTP is used in the event that you want to be able to complete transfers programmatically without use of OTPs
// No arguments required. You will get an OTP to complete the request.
func (c *Client) DisableOTP(ctx context.Context) (*types.Response[any], error) {
	return net.Post[any, any](
		ctx, c.client, c.secret, "/transfer/disable_otp", nil, c.baseURL,
	)
}
