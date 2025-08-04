package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// TransactionTimelineResponse represents the response for viewing transaction timelines
type TransactionTimelineResponse = types.Response[types.Log]

// ViewTimelineByID retrieves the timeline of a transaction by its ID
func (c *Client) ViewTimelineByID(ctx context.Context, id uint64) (*TransactionTimelineResponse, error) {
	return net.Get[types.Log](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%d", basePath, transactionViewTimelinePath, id), "", c.BaseURL)
}

// ViewTimelineByReference retrieves the timeline of a transaction by its reference
func (c *Client) ViewTimelineByReference(ctx context.Context, reference string) (*TransactionTimelineResponse, error) {
	return net.Get[types.Log](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s/%s", basePath, transactionViewTimelinePath, reference), "", c.BaseURL)
}

func (c *Client) ViewTimelineByIDOrReference(ctx context.Context, idOrReference string) (*TransactionTimelineResponse, error) {
	return c.ViewTimelineByReference(ctx, idOrReference)
}
