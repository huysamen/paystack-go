package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckPendingRequest struct {
	Reference string `json:"reference"`
}

type CheckPendingRequestBuilder struct {
	req *CheckPendingRequest
}

func NewCheckPendingChargeRequest(reference string) *CheckPendingRequestBuilder {
	return &CheckPendingRequestBuilder{
		req: &CheckPendingRequest{
			Reference: reference,
		},
	}
}

func (b *CheckPendingRequestBuilder) Build() *CheckPendingRequest {
	return b.req
}

type CheckPendingResponseData = types.ChargeData
type CheckPendingResponse = types.Response[CheckPendingResponseData]

func (c *Client) CheckPending(ctx context.Context, builder CheckPendingRequestBuilder) (*CheckPendingResponse, error) {
	return net.Post[CheckPendingRequest, CheckPendingResponseData](ctx, c.Client, c.Secret, checkPendingPath, builder.Build(), c.BaseURL)
}
