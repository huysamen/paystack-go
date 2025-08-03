package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubmitBirthdayRequest represents the request to submit birthday for a charge
type SubmitBirthdayRequest struct {
	Birthday  string `json:"birthday"`
	Reference string `json:"reference"`
}

// SubmitBirthdayRequestBuilder provides a fluent interface for building SubmitBirthdayRequest
type SubmitBirthdayRequestBuilder struct {
	req *SubmitBirthdayRequest
}

// NewSubmitBirthdayRequest creates a new builder for SubmitBirthdayRequest
func NewSubmitBirthdayRequest(birthday, reference string) *SubmitBirthdayRequestBuilder {
	return &SubmitBirthdayRequestBuilder{
		req: &SubmitBirthdayRequest{
			Birthday:  birthday,
			Reference: reference,
		},
	}
}

// Build returns the constructed SubmitBirthdayRequest
func (b *SubmitBirthdayRequestBuilder) Build() *SubmitBirthdayRequest {
	return b.req
}

type SubmitBirthdayResponse = types.Response[types.ChargeData]

// SubmitBirthday submits birthday when requested for verification
func (c *Client) SubmitBirthday(ctx context.Context, builder *SubmitBirthdayRequestBuilder) (*SubmitBirthdayResponse, error) {
	return net.Post[SubmitBirthdayRequest, types.ChargeData](ctx, c.Client, c.Secret, submitBirthdayPath, builder.Build(), c.BaseURL)
}
