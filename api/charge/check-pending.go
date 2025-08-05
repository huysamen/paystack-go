package charge

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type checkPendingRequest struct {
	Reference string `json:"reference"`
}

type CheckPendingRequestBuilder struct {
	req *checkPendingRequest
}

func NewCheckPendingChargeRequestBuilder(reference string) *CheckPendingRequestBuilder {
	return &CheckPendingRequestBuilder{
		req: &checkPendingRequest{
			Reference: reference,
		},
	}
}

func (b *CheckPendingRequestBuilder) Build() *checkPendingRequest {
	return b.req
}

type CheckPendingResponseData = types.ChargeData
type CheckPendingResponse = types.Response[CheckPendingResponseData]

func (c *Client) CheckPending(ctx context.Context, builder CheckPendingRequestBuilder) (*CheckPendingResponse, error) {
	return net.Post[checkPendingRequest, CheckPendingResponseData](ctx, c.Client, c.Secret, checkPendingPath, builder.Build(), c.BaseURL)
}
