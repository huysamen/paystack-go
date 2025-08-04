package customers

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request type
type CustomerCreateRequest struct {
	Email     string         `json:"email"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	Phone     *string        `json:"phone,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// Builder for creating CustomerCreateRequest
type CustomerCreateRequestBuilder struct {
	req *CustomerCreateRequest
}

// NewCreateCustomerRequest creates a new builder for customer creation
func NewCreateCustomerRequest(email string) *CustomerCreateRequestBuilder {
	return &CustomerCreateRequestBuilder{
		req: &CustomerCreateRequest{
			Email: email,
		},
	}
}

// FirstName sets the first name
func (b *CustomerCreateRequestBuilder) FirstName(firstName string) *CustomerCreateRequestBuilder {
	b.req.FirstName = &firstName

	return b
}

// LastName sets the last name
func (b *CustomerCreateRequestBuilder) LastName(lastName string) *CustomerCreateRequestBuilder {
	b.req.LastName = &lastName

	return b
}

// Phone sets the phone number
func (b *CustomerCreateRequestBuilder) Phone(phone string) *CustomerCreateRequestBuilder {
	b.req.Phone = &phone

	return b
}

// Metadata sets the metadata
func (b *CustomerCreateRequestBuilder) Metadata(metadata map[string]any) *CustomerCreateRequestBuilder {
	b.req.Metadata = metadata

	return b
}

// Build creates the CustomerCreateRequest
func (b *CustomerCreateRequestBuilder) Build() *CustomerCreateRequest {
	return b.req
}

// CustomerCreateResponse represents the response for creating a customer
type CustomerCreateResponse = types.Response[types.Customer]

// Create creates a new customer with the provided builder
func (c *Client) Create(ctx context.Context, builder *CustomerCreateRequestBuilder) (*CustomerCreateResponse, error) {
	return net.Post[CustomerCreateRequest, types.Customer](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
