package bulkcharges

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Resume resumes processing of a paused bulk charge batch
func (c *Client) Resume(ctx context.Context, batchCode string) (*types.Response[any], error) {
	if batchCode == "" {
		return nil, errors.New("batch code is required")
	}

	return net.Get[any](
		ctx, c.client, c.secret, bulkChargesBasePath+"/resume/"+batchCode, c.baseURL,
	)
}
