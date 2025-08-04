package transfercontrol

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type ResendOTPRequest struct {
	TransferCode string `json:"transfer_code"`
	Reason       string `json:"reason"`
}

type ResendOTPRequestBuilder struct {
	request ResendOTPRequest
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

func (b *ResendOTPRequestBuilder) Build() *ResendOTPRequest {
	return &b.request
}

type ResendOTPResponse = types.Response[any]

func (c *Client) ResendOTP(ctx context.Context, builder *ResendOTPRequestBuilder) (*ResendOTPResponse, error) {
	req := builder.Build()
	return net.Post[ResendOTPRequest, any](ctx, c.Client, c.Secret, "/transfer/resend_otp", req, c.BaseURL)
}
