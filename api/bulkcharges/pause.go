package bulkcharges

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Pause pauses processing of a bulk charge batch
func (c *Client) Pause(ctx context.Context, batchCode string) (*types.Response[any], error) {
	if batchCode == "" {
		return nil, errors.New("batch code is required")
	}

	return net.Get[any](
		ctx, c.client, c.secret, bulkChargesBasePath+"/pause/"+batchCode, c.baseURL,
	)
}
