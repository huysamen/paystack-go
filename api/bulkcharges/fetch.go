package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FetchBulkChargeBatch = types.Response[types.BulkChargeBatch]

// Fetch retrieves a specific bulk charge batch by ID or batch code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchBulkChargeBatch, error) {
	return net.Get[types.BulkChargeBatch](ctx, c.Client, c.Secret, basePath+"/"+idOrCode, c.BaseURL)
}
