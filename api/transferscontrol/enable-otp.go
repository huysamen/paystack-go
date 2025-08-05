package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type EnableOTPResponseData = any
type EnableOTPResponse = types.Response[EnableOTPResponseData]

func (c *Client) EnableOTP(ctx context.Context) (*EnableOTPResponse, error) {
	return net.Post[any, EnableOTPResponseData](ctx, c.Client, c.Secret, "/transfer/enable_otp", nil, c.BaseURL)
}
