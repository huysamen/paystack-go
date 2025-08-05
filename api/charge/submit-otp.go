package charge

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type submitOTPRequest struct {
	OTP       string `json:"otp"`
	Reference string `json:"reference"`
}

type SubmitOTPRequestBuilder struct {
	req *submitOTPRequest
}

func NewSubmitOTPRequestBuilder(otp, reference string) *SubmitOTPRequestBuilder {
	return &SubmitOTPRequestBuilder{
		req: &submitOTPRequest{
			OTP:       otp,
			Reference: reference,
		},
	}
}

func (b *SubmitOTPRequestBuilder) Build() *submitOTPRequest {
	return b.req
}

type SubmitOTPResponseData = types.ChargeData
type SubmitOTPResponse = types.Response[SubmitOTPResponseData]

func (c *Client) SubmitOTP(ctx context.Context, builder SubmitOTPRequestBuilder) (*SubmitOTPResponse, error) {
	return net.Post[submitOTPRequest, SubmitOTPResponseData](ctx, c.Client, c.Secret, submitOtpPath, builder.Build(), c.BaseURL)
}
