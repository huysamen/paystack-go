package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchTimeoutResponse represents the response from fetching payment session timeout
type FetchTimeoutResponse = types.Response[FetchTimeoutData]

// FetchTimeoutData contains the payment session timeout information
type FetchTimeoutData struct {
	PaymentSessionTimeout int `json:"payment_session_timeout"`
}

// FetchTimeout retrieves the payment session timeout on your integration
func (c *Client) FetchTimeout(ctx context.Context) (*FetchTimeoutResponse, error) {
	return net.Get[FetchTimeoutData](ctx, c.Client, c.Secret, basePath+"/payment_session_timeout", c.BaseURL)
}
