package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// SubmitAddressRequest represents the request to submit address for a charge
type SubmitAddressRequest struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
	Reference string `json:"reference"`
}

// SubmitAddressRequestBuilder provides a fluent interface for building SubmitAddressRequest
type SubmitAddressRequestBuilder struct {
	req *SubmitAddressRequest
}

// NewSubmitAddressRequest creates a new builder for SubmitAddressRequest
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

// Build returns the constructed SubmitAddressRequest
func (b *SubmitAddressRequestBuilder) Build() *SubmitAddressRequest {
	return b.req
}

type SubmitAddressResponse = types.Response[ChargeData]

// SubmitAddress submits address when requested for verification
func (c *Client) SubmitAddress(ctx context.Context, builder *SubmitAddressRequestBuilder) (*SubmitAddressResponse, error) {
	return net.Post[SubmitAddressRequest, ChargeData](ctx, c.Client, c.Secret, submitAddressPath, builder.Build(), c.BaseURL)
}
