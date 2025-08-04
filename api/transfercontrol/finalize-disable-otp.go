package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// FinalizeDisableOTPRequest represents the request to finalize disabling OTP
type FinalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

// FinalizeDisableOTPRequestBuilder builds a FinalizeDisableOTPRequest
type FinalizeDisableOTPRequestBuilder struct {
	request FinalizeDisableOTPRequest
}

// NewFinalizeDisableOTPRequestBuilder creates a new builder
func NewFinalizeDisableOTPRequestBuilder() *FinalizeDisableOTPRequestBuilder {
	return &FinalizeDisableOTPRequestBuilder{}
}

// OTP sets the OTP for the request
func (b *FinalizeDisableOTPRequestBuilder) OTP(otp string) *FinalizeDisableOTPRequestBuilder {
	b.request.OTP = otp

	return b
}

// Build returns the built FinalizeDisableOTPRequest
func (b *FinalizeDisableOTPRequestBuilder) Build() *FinalizeDisableOTPRequest {
	return &b.request
}

// FinalizeDisableOTPResponse represents the response from finalizing disable OTP
type FinalizeDisableOTPResponse = types.Response[any]

// FinalizeDisableOTP finalizes the request to disable OTP on your transfers
func (c *Client) FinalizeDisableOTP(ctx context.Context, builder *FinalizeDisableOTPRequestBuilder) (*FinalizeDisableOTPResponse, error) {
	req := builder.Build()
	return net.Post[FinalizeDisableOTPRequest, any](ctx, c.Client, c.Secret, "/transfer/disable_otp_finalize", req, c.BaseURL)
}
