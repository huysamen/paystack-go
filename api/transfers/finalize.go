package transfers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request type
type TransferFinalizeRequest struct {
	TransferCode string `json:"transfer_code"` // Transfer code to finalize
	OTP          string `json:"otp"`           // OTP sent to business phone
}

// Builder for creating TransferFinalizeRequest
type TransferFinalizeRequestBuilder struct {
	req *TransferFinalizeRequest
}

// NewFinalizeTransferRequest creates a new builder for transfer finalization
func NewFinalizeTransferRequest(transferCode string, otp string) *TransferFinalizeRequestBuilder {
	return &TransferFinalizeRequestBuilder{
		req: &TransferFinalizeRequest{
			TransferCode: transferCode,
			OTP:          otp,
		},
	}
}

// Build creates the TransferFinalizeRequest
func (b *TransferFinalizeRequestBuilder) Build() *TransferFinalizeRequest {
	return b.req
}

// FinalizeResponse represents the response for finalizing a transfer
type FinalizeResponse = types.Response[types.Transfer]

// Finalize finalizes a transfer with the provided builder
func (c *Client) Finalize(ctx context.Context, builder *TransferFinalizeRequestBuilder) (*FinalizeResponse, error) {
	return net.Post[TransferFinalizeRequest, types.Transfer](ctx, c.Client, c.Secret, basePath+"/finalize_transfer", builder.Build(), c.BaseURL)
}
