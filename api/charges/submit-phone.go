package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
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

// SubmitPhoneResponse represents the response from submitting phone
type SubmitPhoneResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// SubmitPhone submits phone number when requested
func (c *Client) SubmitPhone(ctx context.Context, builder *SubmitPhoneRequestBuilder) (*SubmitPhoneResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	url := c.baseURL + chargesBasePath + "/submit_phone"
	resp, err := net.Post[SubmitPhoneRequest, SubmitPhoneResponse](ctx, c.client, c.secret, url, req)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
