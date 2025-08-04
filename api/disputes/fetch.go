package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchDisputeResponse = types.Response[types.Dispute]

func (c *Client) Fetch(ctx context.Context, disputeID string) (*FetchDisputeResponse, error) {
	return net.Get[types.Dispute](ctx, c.Client, c.Secret, basePath+"/"+disputeID, c.BaseURL)
}
