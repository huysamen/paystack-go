package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Create initiates a payment by integrating multiple payment channels
func (c *Client) Create(ctx context.Context, req *CreateChargeRequest) (*CreateChargeResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("create charge request cannot be nil")
	}

	url := c.baseURL + chargesBasePath
	return net.Post[CreateChargeRequest, ChargeData](ctx, c.client, c.secret, url, req)
}
