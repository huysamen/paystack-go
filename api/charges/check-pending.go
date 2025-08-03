package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CheckPendingChargeRequest represents the request to check a pending charge
type CheckPendingChargeRequest struct {
	Reference string `json:"reference"`
}

// CheckPendingChargeRequestBuilder provides a fluent interface for building CheckPendingChargeRequest
type CheckPendingChargeRequestBuilder struct {
	req *CheckPendingChargeRequest
}

// NewCheckPendingChargeRequest creates a new builder for CheckPendingChargeRequest
func NewCheckPendingChargeRequest(reference string) *CheckPendingChargeRequestBuilder {
	return &CheckPendingChargeRequestBuilder{
		req: &CheckPendingChargeRequest{
			Reference: reference,
		},
	}
}

// Build returns the constructed CheckPendingChargeRequest
func (b *CheckPendingChargeRequestBuilder) Build() *CheckPendingChargeRequest {
	return b.req
}

type CheckPendingChargeResponse = types.Response[types.ChargeData]

// CheckPending checks the status of a pending charge
func (c *Client) CheckPending(ctx context.Context, builder *CheckPendingChargeRequestBuilder) (*CheckPendingChargeResponse, error) {
	return net.Post[CheckPendingChargeRequest, types.ChargeData](ctx, c.Client, c.Secret, checkPendingPath, builder.Build(), c.BaseURL)
}
