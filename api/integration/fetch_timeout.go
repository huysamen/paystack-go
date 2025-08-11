package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchTimeoutResponseData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

type FetchTimeoutResponse = types.Response[FetchTimeoutResponseData]

func (c *Client) FetchTimeout(ctx context.Context) (*FetchTimeoutResponse, error) {
	return net.Get[FetchTimeoutResponseData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", c.BaseURL)
}
