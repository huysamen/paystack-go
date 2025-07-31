package transfers

import (
	"context"
	"errors"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransferFinalizeRequest struct {
	TransferCode string `json:"transfer_code"` // Transfer code to finalize
	OTP          string `json:"otp"`           // OTP sent to business phone
}

func (r *TransferFinalizeRequest) Validate() error {
	if r.TransferCode == "" {
		return errors.New("transfer_code is required")
	}
	if r.OTP == "" {
		return errors.New("otp is required")
	}
	return nil
}

func (c *Client) Finalize(ctx context.Context, req *TransferFinalizeRequest) (*types.Response[Transfer], error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	path := transferBasePath + "/finalize_transfer"

	return net.Post[TransferFinalizeRequest, Transfer](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
