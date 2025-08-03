package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
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

type SubmitOTPResponse = types.Response[ChargeData]

// SubmitOTP submits OTP to complete a charge
func (c *Client) SubmitOTP(ctx context.Context, builder *SubmitOTPRequestBuilder) (*SubmitOTPResponse, error) {
	return net.Post[SubmitOTPRequest, ChargeData](ctx, c.Client, c.Secret, submitOtpPath, builder.Build(), c.BaseURL)
}
