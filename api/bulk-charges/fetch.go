package bulkcharges

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Fetch retrieves a specific bulk charge batch by ID or batch code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*types.Response[BulkChargeBatch], error) {
	if idOrCode == "" {
		return nil, errors.New("bulk charge batch ID or code is required")
	}

	return net.Get[BulkChargeBatch](
		ctx, c.client, c.secret, bulkChargesBasePath+"/"+idOrCode, c.baseURL,
	)
}
