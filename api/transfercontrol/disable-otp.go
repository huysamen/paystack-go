package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DisableOTPResponse = types.Response[any]

func (c *Client) DisableOTP(ctx context.Context) (*DisableOTPResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, "/transfer/disable_otp", nil, c.BaseURL)
}
