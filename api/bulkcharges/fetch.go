package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific bulk charge batch by ID or batch code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[BulkChargeBatch], error) {
	return net.Get[BulkChargeBatch](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, c.BaseURL)
}
