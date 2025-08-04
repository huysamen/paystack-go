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

type SubmitBirthdayResponseData = types.ChargeData
type SubmitBirthdayResponse = types.Response[SubmitBirthdayResponseData]

func (c *Client) SubmitBirthday(ctx context.Context, builder SubmitBirthdayRequestBuilder) (*SubmitBirthdayResponse, error) {
	return net.Post[SubmitBirthdayRequest, SubmitBirthdayResponseData](ctx, c.Client, c.Secret, submitBirthdayPath, builder.Build(), c.BaseURL)
}
