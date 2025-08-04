package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type EnableOTPResponse = types.Response[any]

func (c *Client) EnableOTP(ctx context.Context) (*EnableOTPResponse, error) {
	return net.Post[any, any](ctx, c.Client, c.Secret, "/transfer/enable_otp", nil, c.BaseURL)
}
