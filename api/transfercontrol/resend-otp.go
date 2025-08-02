package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// ResendOTPRequestBuilder builds a ResendOTPRequest
type ResendOTPRequestBuilder struct {
	request ResendOTPRequest
}

// NewResendOTPRequestBuilder creates a new builder
func NewResendOTPRequestBuilder() *ResendOTPRequestBuilder {
	return &ResendOTPRequestBuilder{}
}

// TransferCode sets the transfer code for the request
func (b *ResendOTPRequestBuilder) TransferCode(transferCode string) *ResendOTPRequestBuilder {
	b.request.TransferCode = transferCode
	return b
}

// Reason sets the reason for the request
func (b *ResendOTPRequestBuilder) Reason(reason string) *ResendOTPRequestBuilder {
	b.request.Reason = reason
	return b
}

// Build returns the built ResendOTPRequest
func (b *ResendOTPRequestBuilder) Build() *ResendOTPRequest {
	return &b.request
}

// ResendOTP generates a new OTP and sends to customer in the event they are having trouble receiving one
func (c *Client) ResendOTP(ctx context.Context, builder *ResendOTPRequestBuilder) (*types.Response[any], error) {
	req := builder.Build()
	return net.Post[ResendOTPRequest, any](ctx, c.Client, c.Secret, "/transfer/resend_otp", req, c.BaseURL)
}
