package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request type
type CustomerUpdateRequest struct {
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// Builder for updating CustomerUpdateRequest
type CustomerUpdateRequestBuilder struct {
	req *CustomerUpdateRequest
}

// NewUpdateCustomerRequest creates a new builder for customer update
func NewUpdateCustomerRequest() *CustomerUpdateRequestBuilder {
	return &CustomerUpdateRequestBuilder{
		req: &CustomerUpdateRequest{},
	}
}

// FirstName sets the first name
func (b *CustomerUpdateRequestBuilder) FirstName(firstName string) *CustomerUpdateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

// LastName sets the last name
func (b *CustomerUpdateRequestBuilder) LastName(lastName string) *CustomerUpdateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

// Phone sets the phone number
func (b *CustomerUpdateRequestBuilder) Phone(phone string) *CustomerUpdateRequestBuilder {
	b.req.Phone = &phone

	return b
}

// Metadata sets the metadata
func (b *CustomerUpdateRequestBuilder) Metadata(metadata map[string]any) *CustomerUpdateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

// Build creates the CustomerUpdateRequest
func (b *CustomerUpdateRequestBuilder) Build() *CustomerUpdateRequest {
	return b.req
}

// UpdateCustomerResponse represents the response for updating a customer
type UpdateCustomerResponse = types.Response[types.Customer]

// Update updates a customer with the provided builder
func (c *Client) Update(ctx context.Context, customerCode string, builder *CustomerUpdateRequestBuilder) (*UpdateCustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", basePath, customerCode)

	return net.Put[CustomerUpdateRequest, types.Customer](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
