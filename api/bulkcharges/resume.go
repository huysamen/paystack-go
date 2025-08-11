package bulkcharges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ResumeResponseData = any
type ResumeResponse = types.Response[ResumeResponseData]

func (c *Client) Resume(ctx context.Context, batchCode string) (*ResumeResponse, error) {
	return net.Get[ResumeResponseData](ctx, c.Client, c.Secret, resumePath+"/"+batchCode, c.BaseURL)
}
