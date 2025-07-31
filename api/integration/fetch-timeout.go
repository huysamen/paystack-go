package integration

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// FetchTimeout retrieves the payment session timeout on your integration
func (c *Client) FetchTimeout(ctx context.Context) (*FetchTimeoutResponse, error) {
	url := c.baseURL + integrationBasePath + "/payment_session_timeout"
	return net.Get[FetchTimeoutData](ctx, c.client, c.secret, url)
}
