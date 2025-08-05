package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type submitPhoneRequest struct {
	Phone     string `json:"phone"`
	Reference string `json:"reference"`
}

type SubmitPhoneRequestBuilder struct {
	req *submitPhoneRequest
}

func NewSubmitPhoneRequestBuilder(phone, reference string) *SubmitPhoneRequestBuilder {
	return &SubmitPhoneRequestBuilder{
		req: &submitPhoneRequest{
			Phone:     phone,
			Reference: reference,
		},
	}
}

func (b *SubmitPhoneRequestBuilder) Build() *submitPhoneRequest {
	return b.req
}

type SubmitPhoneResponseData = types.ChargeData
type SubmitPhoneResponse = types.Response[SubmitPhoneResponseData]

func (c *Client) SubmitPhone(ctx context.Context, builder SubmitPhoneRequestBuilder) (*SubmitPhoneResponse, error) {
	return net.Post[submitPhoneRequest, SubmitPhoneResponseData](ctx, c.Client, c.Secret, submitPhonePath, builder.Build(), c.BaseURL)
}
