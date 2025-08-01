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
	resp, err := net.Get[FetchTimeoutData](
		ctx,
		c.client,
		c.secret,
		integrationBasePath+"/payment_session_timeout",
		c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
