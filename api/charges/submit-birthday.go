package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubmitBirthdayRequest struct {
	Birthday  string `json:"birthday"`
	Reference string `json:"reference"`
}

type SubmitBirthdayRequestBuilder struct {
	req *SubmitBirthdayRequest
}

func NewSubmitBirthdayRequest(birthday, reference string) *SubmitBirthdayRequestBuilder {
	return &SubmitBirthdayRequestBuilder{
		req: &SubmitBirthdayRequest{
			Birthday:  birthday,
			Reference: reference,
		},
	}
}

func (b *SubmitBirthdayRequestBuilder) Build() *SubmitBirthdayRequest {
	return b.req
}

type SubmitBirthdayResponse = types.Response[types.ChargeData]

func (c *Client) SubmitBirthday(ctx context.Context, builder *SubmitBirthdayRequestBuilder) (*SubmitBirthdayResponse, error) {
	return net.Post[SubmitBirthdayRequest, types.ChargeData](ctx, c.Client, c.Secret, submitBirthdayPath, builder.Build(), c.BaseURL)
}
