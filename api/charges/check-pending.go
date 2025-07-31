package charges

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// CheckPending checks the status of a pending charge
func (c *Client) CheckPending(ctx context.Context, reference string) (*CheckPendingChargeResponse, error) {
	if reference == "" {
		return nil, fmt.Errorf("reference is required")
	}

	url := c.baseURL + chargesBasePath + "/" + reference
	return net.Get[ChargeData](ctx, c.client, c.secret, url)
}
