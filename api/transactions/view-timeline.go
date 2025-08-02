package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

func (c *Client) ViewTimelineByID(ctx context.Context, id uint64) (*types.Response[types.Log], error) {
	return net.Get[types.Log](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%d", basePath, transactionViewTimelinePath, id), "", c.BaseURL)
}

func (c *Client) ViewTimelineByReference(ctx context.Context, reference string) (*types.Response[types.Log], error) {
	return net.Get[types.Log](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%s", basePath, transactionViewTimelinePath, reference), "", c.BaseURL)
}
