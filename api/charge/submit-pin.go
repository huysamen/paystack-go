package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type submitPINRequest struct {
	PIN       string `json:"pin"`
	Reference string `json:"reference"`
}

type SubmitPINRequestBuilder struct {
	req *submitPINRequest
}

func NewSubmitPINRequestBuilder(pin, reference string) *SubmitPINRequestBuilder {
	return &SubmitPINRequestBuilder{
		req: &submitPINRequest{
			PIN:       pin,
			Reference: reference,
		},
	}
}

func (b *SubmitPINRequestBuilder) Build() *submitPINRequest {
	return b.req
}

type SubmitPINResponseData = types.ChargeData
type SubmitPINResponse = types.Response[SubmitPINResponseData]

func (c *Client) SubmitPIN(ctx context.Context, builder SubmitPINRequestBuilder) (*SubmitPINResponse, error) {
	return net.Post[submitPINRequest, SubmitPINResponseData](ctx, c.Client, c.Secret, submitPinPath, builder.Build(), c.BaseURL)
}
