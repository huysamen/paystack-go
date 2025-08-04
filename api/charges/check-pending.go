package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CheckPendingChargeRequest struct {
	Reference string `json:"reference"`
}

type CheckPendingChargeRequestBuilder struct {
	req *CheckPendingChargeRequest
}

func NewCheckPendingChargeRequest(reference string) *CheckPendingChargeRequestBuilder {
	return &CheckPendingChargeRequestBuilder{
		req: &CheckPendingChargeRequest{
			Reference: reference,
		},
	}
}

func (b *CheckPendingChargeRequestBuilder) Build() *CheckPendingChargeRequest {
	return b.req
}

type CheckPendingChargeResponse = types.Response[types.ChargeData]

func (c *Client) CheckPending(ctx context.Context, builder *CheckPendingChargeRequestBuilder) (*CheckPendingChargeResponse, error) {
	return net.Post[CheckPendingChargeRequest, types.ChargeData](ctx, c.Client, c.Secret, checkPendingPath, builder.Build(), c.BaseURL)
}
