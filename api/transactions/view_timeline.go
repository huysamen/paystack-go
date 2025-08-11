package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TimelineResponseData = types.TransactionLog
type TimelineResponse = types.Response[TimelineResponseData]

func (c *Client) ViewTimelineByID(ctx context.Context, id uint64) (*TimelineResponse, error) {
	return net.Get[TimelineResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%d", basePath, transactionViewTimelinePath, id), "", c.BaseURL)
}

func (c *Client) ViewTimelineByReference(ctx context.Context, reference string) (*TimelineResponse, error) {
	return net.Get[TimelineResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%s", basePath, transactionViewTimelinePath, reference), "", c.BaseURL)
}

func (c *Client) ViewTimelineByIDOrReference(ctx context.Context, idOrReference string) (*TimelineResponse, error) {
	return c.ViewTimelineByReference(ctx, idOrReference)
}
