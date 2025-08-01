package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
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

// SubmitBirthdayResponse represents the response from submitting birthday
type SubmitBirthdayResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// SubmitBirthday submits birthday when requested
func (c *Client) SubmitBirthday(ctx context.Context, builder *SubmitBirthdayRequestBuilder) (*SubmitBirthdayResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	url := c.baseURL + chargesBasePath + "/submit_birthday"
	resp, err := net.Post[SubmitBirthdayRequest, SubmitBirthdayResponse](ctx, c.client, c.secret, url, req)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
