package charge

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type submitBirthdayRequest struct {
	Birthday  string `json:"birthday"`
	Reference string `json:"reference"`
}

type SubmitBirthdayRequestBuilder struct {
	req *submitBirthdayRequest
}

func NewSubmitBirthdayRequestBuilder(birthday, reference string) *SubmitBirthdayRequestBuilder {
	return &SubmitBirthdayRequestBuilder{
		req: &submitBirthdayRequest{
			Birthday:  birthday,
			Reference: reference,
		},
	}
}

func (b *SubmitBirthdayRequestBuilder) Build() *submitBirthdayRequest {
	return b.req
}

type SubmitBirthdayResponseData = types.Charge
type SubmitBirthdayResponse = types.Response[SubmitBirthdayResponseData]

func (c *Client) SubmitBirthday(ctx context.Context, builder SubmitBirthdayRequestBuilder) (*SubmitBirthdayResponse, error) {
	return net.Post[submitBirthdayRequest, SubmitBirthdayResponseData](ctx, c.Client, c.Secret, submitBirthdayPath, builder.Build(), c.BaseURL)
}
