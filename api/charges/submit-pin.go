package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
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

// SubmitPINResponse represents the response from submitting PIN
type SubmitPINResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// SubmitPIN submits PIN to continue a charge
func (c *Client) SubmitPIN(ctx context.Context, builder *SubmitPINRequestBuilder) (*SubmitPINResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	url := c.baseURL + chargesBasePath + "/submit_pin"
	resp, err := net.Post[SubmitPINRequest, SubmitPINResponse](ctx, c.client, c.secret, url, req)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
