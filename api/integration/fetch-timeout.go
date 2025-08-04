package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchTimeoutResponse = types.Response[FetchTimeoutData]

type FetchTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

func (c *Client) FetchTimeout(ctx context.Context) (*FetchTimeoutResponse, error) {
	return net.Get[FetchTimeoutData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", c.BaseURL)
}
