package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FetchDisputeResponse represents the response from fetching a dispute
type FetchDisputeResponse = types.Response[Dispute]

// Fetch retrieves details of a specific dispute
func (c *Client) Fetch(ctx context.Context, disputeID string) (*FetchDisputeResponse, error) {
	return net.Get[Dispute](ctx, c.Client, c.Secret, basePath+"/"+disputeID, c.BaseURL)
}
