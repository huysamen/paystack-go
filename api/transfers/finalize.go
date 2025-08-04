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

func (c *Client) Finalize(ctx context.Context, req *TransferFinalizeRequest) (*types.Response[types.Transfer], error) {
	return net.Post[TransferFinalizeRequest, types.Transfer](ctx, c.Client, c.Secret, basePath+"/finalize_transfer", req, c.BaseURL)
}
