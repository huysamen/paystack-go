package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type SubmitAddressRequest struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
	Reference string `json:"reference"`
}

type SubmitAddressRequestBuilder struct {
	req *SubmitAddressRequest
}

func NewSubmitAddressRequest(address, city, state, zipCode, reference string) *SubmitAddressRequestBuilder {
	return &SubmitAddressRequestBuilder{
		req: &SubmitAddressRequest{
			Address:   address,
			City:      city,
			State:     state,
			ZipCode:   zipCode,
			Reference: reference,
		},
	}
}

func (b *SubmitAddressRequestBuilder) Build() *SubmitAddressRequest {
	return b.req
}

type SubmitAddressResponse = types.Response[types.ChargeData]

func (c *Client) SubmitAddress(ctx context.Context, builder *SubmitAddressRequestBuilder) (*SubmitAddressResponse, error) {
	return net.Post[SubmitAddressRequest, types.ChargeData](ctx, c.Client, c.Secret, submitAddressPath, builder.Build(), c.BaseURL)
}
