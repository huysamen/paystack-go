package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubmitPINRequest struct {
	PIN       string `json:"pin"`
	Reference string `json:"reference"`
}

type SubmitPINRequestBuilder struct {
	req *SubmitPINRequest
}

func NewSubmitPINRequest(pin, reference string) *SubmitPINRequestBuilder {
	return &SubmitPINRequestBuilder{
		req: &SubmitPINRequest{
			PIN:       pin,
			Reference: reference,
		},
	}
}

func (b *SubmitPINRequestBuilder) Build() *SubmitPINRequest {
	return b.req
}

type SubmitPINResponseData = types.ChargeData
type SubmitPINResponse = types.Response[SubmitPINResponseData]

func (c *Client) SubmitPIN(ctx context.Context, builder SubmitPINRequestBuilder) (*SubmitPINResponse, error) {
	return net.Post[SubmitPINRequest, SubmitPINResponseData](ctx, c.Client, c.Secret, submitPinPath, builder.Build(), c.BaseURL)
}
