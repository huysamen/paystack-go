package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubmitPINRequest represents the request to submit PIN for a charge
type SubmitPINRequest struct {
	PIN       string `json:"pin"`
	Reference string `json:"reference"`
}

// SubmitPINRequestBuilder provides a fluent interface for building SubmitPINRequest
type SubmitPINRequestBuilder struct {
	req *SubmitPINRequest
}

// NewSubmitPINRequest creates a new builder for SubmitPINRequest
func NewSubmitPINRequest(pin, reference string) *SubmitPINRequestBuilder {
	return &SubmitPINRequestBuilder{
		req: &SubmitPINRequest{
			PIN:       pin,
			Reference: reference,
		},
	}
}

// Build returns the constructed SubmitPINRequest
func (b *SubmitPINRequestBuilder) Build() *SubmitPINRequest {
	return b.req
}

type SubmitPINResponse = types.Response[ChargeData]

// SubmitPIN submits PIN to continue a charge
func (c *Client) SubmitPIN(ctx context.Context, builder *SubmitPINRequestBuilder) (*SubmitPINResponse, error) {
	return net.Post[SubmitPINRequest, ChargeData](ctx, c.Client, c.Secret, submitPinPath, builder.Build(), c.BaseURL)
}
