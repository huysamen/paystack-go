package integration

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// UpdateTimeout updates the payment session timeout on your integration
func (c *Client) UpdateTimeout(ctx context.Context, req *UpdateTimeoutRequest) (*UpdateTimeoutResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("update timeout request cannot be nil")
	}

	url := c.baseURL + integrationBasePath + "/payment_session_timeout"
	return net.Put[UpdateTimeoutRequest, UpdateTimeoutData](ctx, c.client, c.secret, url, req)
}
