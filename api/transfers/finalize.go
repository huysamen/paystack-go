package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type finalizeRequest struct {
	TransferCode string `json:"transfer_code"` // Transfer code to finalize
	OTP          string `json:"otp"`           // OTP sent to business phone
}

type FinalizeRequestBuilder struct {
	req *finalizeRequest
}

func NewFinalizeRequestBuilder(transferCode string, otp string) *FinalizeRequestBuilder {
	return &FinalizeRequestBuilder{
		req: &finalizeRequest{
			TransferCode: transferCode,
			OTP:          otp,
		},
	}
}

func (b *FinalizeRequestBuilder) Build() *finalizeRequest {
	return b.req
}

type FinalizeResponseData = types.Transfer
type FinalizeResponse = types.Response[FinalizeResponseData]

func (c *Client) Finalize(ctx context.Context, builder FinalizeRequestBuilder) (*FinalizeResponse, error) {
	return net.Post[finalizeRequest, FinalizeResponseData](ctx, c.Client, c.Secret, basePath+"/finalize_transfer", builder.Build(), c.BaseURL)
}
