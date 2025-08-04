package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type submitAddressRequest struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
	Reference string `json:"reference"`
}

type SubmitAddressRequestBuilder struct {
	req *submitAddressRequest
}

func NewSubmitAddressRequestBuilder(address, city, state, zipCode, reference string) *SubmitAddressRequestBuilder {
	return &SubmitAddressRequestBuilder{
		req: &submitAddressRequest{
			Address:   address,
			City:      city,
			State:     state,
			ZipCode:   zipCode,
			Reference: reference,
		},
	}
}

func (b *SubmitAddressRequestBuilder) Build() *submitAddressRequest {
	return b.req
}

type SubmitAddressResponseData = types.ChargeData
type SubmitAddressResponse = types.Response[SubmitAddressResponseData]

func (c *Client) SubmitAddress(ctx context.Context, builder SubmitAddressRequestBuilder) (*SubmitAddressResponse, error) {
	return net.Post[submitAddressRequest, SubmitAddressResponseData](ctx, c.Client, c.Secret, submitAddressPath, builder.Build(), c.BaseURL)
}
