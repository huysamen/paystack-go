package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Pause pauses processing of a bulk charge batch
func (c *Client) Pause(ctx context.Context, batchCode string) (*types.Response[any], error) {
	return net.Get[any](ctx, c.Client, c.Secret, pausePath+"/"+batchCode, c.BaseURL)
}
