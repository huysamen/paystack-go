package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferFinalizeRequest struct {
	TransferCode string `json:"transfer_code"` // Transfer code to finalize
	OTP          string `json:"otp"`           // OTP sent to business phone
}

type TransferFinalizeRequestBuilder struct {
	req *TransferFinalizeRequest
}

func NewFinalizeTransferRequest(transferCode string, otp string) *TransferFinalizeRequestBuilder {
	return &TransferFinalizeRequestBuilder{
		req: &TransferFinalizeRequest{
			TransferCode: transferCode,
			OTP:          otp,
		},
	}
}

func (b *TransferFinalizeRequestBuilder) Build() *TransferFinalizeRequest {
	return b.req
}

type FinalizeResponse = types.Response[types.Transfer]

func (c *Client) Finalize(ctx context.Context, builder *TransferFinalizeRequestBuilder) (*FinalizeResponse, error) {
	return net.Post[TransferFinalizeRequest, types.Transfer](ctx, c.Client, c.Secret, basePath+"/finalize_transfer", builder.Build(), c.BaseURL)
}
