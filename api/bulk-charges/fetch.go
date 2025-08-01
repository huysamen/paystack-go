package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// FetchBulkChargeBatchResponse represents the response from fetching a bulk charge batch
type FetchBulkChargeBatchResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    BulkChargeBatch `json:"data"`
}

// Fetch retrieves a specific bulk charge batch by ID or batch code
func (c *Client) Fetch(ctx context.Context, idOrCode string) (*FetchBulkChargeBatchResponse, error) {
	resp, err := net.Get[FetchBulkChargeBatchResponse](
		ctx, c.client, c.secret, bulkChargesBasePath+"/"+idOrCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
