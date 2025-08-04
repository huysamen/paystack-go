package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type FinalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

type FinalizeDisableOTPRequestBuilder struct {
	request FinalizeDisableOTPRequest
}

func NewFinalizeDisableOTPRequestBuilder() *FinalizeDisableOTPRequestBuilder {
	return &FinalizeDisableOTPRequestBuilder{}
}

func (b *FinalizeDisableOTPRequestBuilder) OTP(otp string) *FinalizeDisableOTPRequestBuilder {
	b.request.OTP = otp

	return b
}

func (b *FinalizeDisableOTPRequestBuilder) Build() *FinalizeDisableOTPRequest {
	return &b.request
}

type FinalizeDisableOTPResponse = types.Response[any]

func (c *Client) FinalizeDisableOTP(ctx context.Context, builder *FinalizeDisableOTPRequestBuilder) (*FinalizeDisableOTPResponse, error) {
	req := builder.Build()
	return net.Post[FinalizeDisableOTPRequest, any](ctx, c.Client, c.Secret, "/transfer/disable_otp_finalize", req, c.BaseURL)
}
