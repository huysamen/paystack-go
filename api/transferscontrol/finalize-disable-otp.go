package transferscontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type finalizeDisableOTPRequest struct {
	OTP string `json:"otp"`
}

type FinalizeDisableOTPRequestBuilder struct {
	request finalizeDisableOTPRequest
}

func NewFinalizeDisableOTPRequestBuilder() *FinalizeDisableOTPRequestBuilder {
	return &FinalizeDisableOTPRequestBuilder{}
}

func (b *FinalizeDisableOTPRequestBuilder) OTP(otp string) *FinalizeDisableOTPRequestBuilder {
	b.request.OTP = otp

	return b
}

func (b *FinalizeDisableOTPRequestBuilder) Build() *finalizeDisableOTPRequest {
	return &b.request
}

type FinalizeDisableOTPResponseData = any
type FinalizeDisableOTPResponse = types.Response[FinalizeDisableOTPResponseData]

func (c *Client) FinalizeDisableOTP(ctx context.Context, builder FinalizeDisableOTPRequestBuilder) (*FinalizeDisableOTPResponse, error) {
	req := builder.Build()
	return net.Post[finalizeDisableOTPRequest, FinalizeDisableOTPResponseData](ctx, c.Client, c.Secret, "/transfer/disable_otp_finalize", req, c.BaseURL)
}
