package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// PauseBulkChargeBatchResponse represents the response from pausing a bulk charge batch
type PauseBulkChargeBatchResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Pause pauses processing of a bulk charge batch
func (c *Client) Pause(ctx context.Context, batchCode string) (*PauseBulkChargeBatchResponse, error) {
	resp, err := net.Get[PauseBulkChargeBatchResponse](
		ctx, c.client, c.secret, bulkChargesBasePath+"/pause/"+batchCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
