package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubmitOTPRequest struct {
	OTP       string `json:"otp"`
	Reference string `json:"reference"`
}

type SubmitOTPRequestBuilder struct {
	req *SubmitOTPRequest
}

func NewSubmitOTPRequest(otp, reference string) *SubmitOTPRequestBuilder {
	return &SubmitOTPRequestBuilder{
		req: &SubmitOTPRequest{
			OTP:       otp,
			Reference: reference,
		},
	}
}

func (b *SubmitOTPRequestBuilder) Build() *SubmitOTPRequest {
	return b.req
}

type SubmitOTPResponse = types.Response[types.ChargeData]

func (c *Client) SubmitOTP(ctx context.Context, builder *SubmitOTPRequestBuilder) (*SubmitOTPResponse, error) {
	return net.Post[SubmitOTPRequest, types.ChargeData](ctx, c.Client, c.Secret, submitOtpPath, builder.Build(), c.BaseURL)
}
