package disputes

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchResponseData = types.Dispute
type FetchResponse = types.Response[FetchResponseData]

func (c *Client) Fetch(ctx context.Context, disputeID string) (*FetchResponse, error) {
	return net.Get[FetchResponseData](ctx, c.Client, c.Secret, basePath+"/"+disputeID, c.BaseURL)
}
