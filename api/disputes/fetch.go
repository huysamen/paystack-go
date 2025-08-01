package disputes

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
)

// Fetch retrieves details of a specific dispute
func (c *Client) Fetch(ctx context.Context, disputeID string) (*DisputeFetchResponse, error) {
	if disputeID == "" {
		return nil, fmt.Errorf("dispute ID is required")
	}

	url := c.baseURL + disputesBasePath + "/" + disputeID
	return net.Get[Dispute](ctx, c.client, c.secret, url)
}
