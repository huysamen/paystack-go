package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubmitPhoneRequest represents the request to submit phone number for a charge
type SubmitPhoneRequest struct {
	Phone     string `json:"phone"`
	Reference string `json:"reference"`
}

// SubmitPhoneRequestBuilder provides a fluent interface for building SubmitPhoneRequest
type SubmitPhoneRequestBuilder struct {
	req *SubmitPhoneRequest
}

// NewSubmitPhoneRequest creates a new builder for SubmitPhoneRequest
func NewSubmitPhoneRequest(phone, reference string) *SubmitPhoneRequestBuilder {
	return &SubmitPhoneRequestBuilder{
		req: &SubmitPhoneRequest{
			Phone:     phone,
			Reference: reference,
		},
	}
}

// Build returns the constructed SubmitPhoneRequest
func (b *SubmitPhoneRequestBuilder) Build() *SubmitPhoneRequest {
	return b.req
}

// SubmitPhone submits phone number when requested
func (c *Client) SubmitPhone(ctx context.Context, builder *SubmitPhoneRequestBuilder) (*types.Response[ChargeData], error) {
	return net.Post[SubmitPhoneRequest, ChargeData](ctx, c.Client, c.Secret, submitPhonePath, builder.Build(), c.BaseURL)
}
