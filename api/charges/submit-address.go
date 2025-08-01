package charges

import (
	"context"

	"github.com/huysamen/paystack-go/net"
)

// SubmitAddressRequest represents the request to submit address for a charge
type SubmitAddressRequest struct {
	Address   string `json:"address"`
	Reference string `json:"reference"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}

// SubmitAddressRequestBuilder provides a fluent interface for building SubmitAddressRequest
type SubmitAddressRequestBuilder struct {
	req *SubmitAddressRequest
}

// NewSubmitAddressRequest creates a new builder for SubmitAddressRequest
func NewSubmitAddressRequest(address, reference, city, state, zipCode string) *SubmitAddressRequestBuilder {
	return &SubmitAddressRequestBuilder{
		req: &SubmitAddressRequest{
			Address:   address,
			Reference: reference,
			City:      city,
			State:     state,
			ZipCode:   zipCode,
		},
	}
}

// Build returns the constructed SubmitAddressRequest
func (b *SubmitAddressRequestBuilder) Build() *SubmitAddressRequest {
	return b.req
}

// SubmitAddressResponse represents the response from submitting address
type SubmitAddressResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    ChargeData `json:"data"`
}

// SubmitAddress submits address to continue a charge
func (c *Client) SubmitAddress(ctx context.Context, builder *SubmitAddressRequestBuilder) (*SubmitAddressResponse, error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	url := c.baseURL + chargesBasePath + "/submit_address"
	resp, err := net.Post[SubmitAddressRequest, SubmitAddressResponse](ctx, c.client, c.secret, url, req)
	if err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
