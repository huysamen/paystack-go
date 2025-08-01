package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// SubmitOTPRequest represents the request to submit OTP for a charge
type SubmitOTPRequest struct {
	OTP       string `json:"otp"`
	Reference string `json:"reference"`
}

// SubmitOTPRequestBuilder provides a fluent interface for building SubmitOTPRequest
type SubmitOTPRequestBuilder struct {
	req *SubmitOTPRequest
}

// NewSubmitOTPRequest creates a new builder for SubmitOTPRequest
func NewSubmitOTPRequest(otp, reference string) *SubmitOTPRequestBuilder {
	return &SubmitOTPRequestBuilder{
		req: &SubmitOTPRequest{
			OTP:       otp,
			Reference: reference,
		},
	}
}

// Build returns the constructed SubmitOTPRequest
func (b *SubmitOTPRequestBuilder) Build() *SubmitOTPRequest {
	return b.req
}

// SubmitOTPResponse represents the response from submitting OTP
type SubmitOTPResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// SubmitOTP submits OTP to complete a charge
func (c *Client) SubmitOTP(ctx context.Context, builder *SubmitOTPRequestBuilder) (*SubmitOTPResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	url := c.baseURL + chargesBasePath + "/submit_otp"
	resp, err := net.Post[SubmitOTPRequest, SubmitOTPResponse](ctx, c.client, c.secret, url, req)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
