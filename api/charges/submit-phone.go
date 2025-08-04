package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubmitPhoneRequest struct {
	Phone     string `json:"phone"`
	Reference string `json:"reference"`
}

type SubmitPhoneRequestBuilder struct {
	req *SubmitPhoneRequest
}

func NewSubmitPhoneRequest(phone, reference string) *SubmitPhoneRequestBuilder {
	return &SubmitPhoneRequestBuilder{
		req: &SubmitPhoneRequest{
			Phone:     phone,
			Reference: reference,
		},
	}
}

func (b *SubmitPhoneRequestBuilder) Build() *SubmitPhoneRequest {
	return b.req
}

type SubmitPhoneResponseData = types.ChargeData
type SubmitPhoneResponse = types.Response[SubmitPhoneResponseData]

func (c *Client) SubmitPhone(ctx context.Context, builder SubmitPhoneRequestBuilder) (*SubmitPhoneResponse, error) {
	return net.Post[SubmitPhoneRequest, SubmitPhoneResponseData](ctx, c.Client, c.Secret, submitPhonePath, builder.Build(), c.BaseURL)
}
