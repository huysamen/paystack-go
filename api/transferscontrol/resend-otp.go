package transferscontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type resendOTPRequest struct {
	TransferCode string `json:"transfer_code"`
	Reason       string `json:"reason"`
}

type ResendOTPRequestBuilder struct {
	request resendOTPRequest
}

func NewResendOTPRequestBuilder() *ResendOTPRequestBuilder {
	return &ResendOTPRequestBuilder{}
}

func (b *ResendOTPRequestBuilder) TransferCode(transferCode string) *ResendOTPRequestBuilder {
	b.request.TransferCode = transferCode

	return b
}

func (b *ResendOTPRequestBuilder) Reason(reason string) *ResendOTPRequestBuilder {
	b.request.Reason = reason

	return b
}

func (b *ResendOTPRequestBuilder) Build() *resendOTPRequest {
	return &b.request
}

type ResendOTPResponseData = any
type ResendOTPResponse = types.Response[ResendOTPResponseData]

func (c *Client) ResendOTP(ctx context.Context, builder ResendOTPRequestBuilder) (*ResendOTPResponse, error) {
	return net.Post[resendOTPRequest, ResendOTPResponseData](ctx, c.Client, c.Secret, "/transfer/resend_otp", builder.Build(), c.BaseURL)
}
