package transferscontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type DisableOTPResponseData = any
type DisableOTPResponse = types.Response[DisableOTPResponseData]

func (c *Client) DisableOTP(ctx context.Context) (*DisableOTPResponse, error) {
	return net.Post[any, DisableOTPResponseData](ctx, c.Client, c.Secret, "/transfer/disable_otp", nil, c.BaseURL)
}
