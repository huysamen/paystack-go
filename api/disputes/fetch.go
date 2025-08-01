package disputes

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchDisputeResponse represents the response from fetching a dispute
type FetchDisputeResponse = types.Response[Dispute]

// Fetch retrieves details of a specific dispute
func (c *Client) Fetch(ctx context.Context, disputeID string) (*types.Response[Dispute], error) {
	if disputeID == "" {
		return nil, errors.New("dispute ID is required")
	}

	endpoint := c.baseURL + disputesBasePath + "/" + disputeID
	resp, err := net.Get[Dispute](ctx, c.client, c.secret, endpoint, c.baseURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
