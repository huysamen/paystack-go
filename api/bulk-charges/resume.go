package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// ResumeBulkChargeBatchResponse represents the response from resuming a bulk charge batch
type ResumeBulkChargeBatchResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// Resume resumes processing of a paused bulk charge batch
func (c *Client) Resume(ctx context.Context, batchCode string) (*ResumeBulkChargeBatchResponse, error) {
	resp, err := net.Get[ResumeBulkChargeBatchResponse](
		ctx, c.client, c.secret, bulkChargesBasePath+"/resume/"+batchCode, c.baseURL,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
