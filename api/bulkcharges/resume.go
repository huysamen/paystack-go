package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ResumeBulkChargeResponse = types.Response[any]

// Resume resumes processing of a paused bulk charge batch
func (c *Client) Resume(ctx context.Context, batchCode string) (*ResumeBulkChargeResponse, error) {
	return net.Get[any](ctx, c.Client, c.Secret, resumePath+"/"+batchCode, c.BaseURL)
}
