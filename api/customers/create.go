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
	email     string
	firstName *string
	lastName  *string
	phone     *string
	metadata  map[string]any
}

// NewCreateCustomerRequest creates a new builder for customer creation
func NewCreateCustomerRequest(email string) *CustomerCreateRequestBuilder {
	return &CustomerCreateRequestBuilder{
		email: email,
	}
}

// FirstName sets the first name
func (b *CustomerCreateRequestBuilder) FirstName(firstName string) *CustomerCreateRequestBuilder {
	b.firstName = &firstName
	return b
}

// LastName sets the last name
func (b *CustomerCreateRequestBuilder) LastName(lastName string) *CustomerCreateRequestBuilder {
	b.lastName = &lastName
	return b
}

// Phone sets the phone number
func (b *CustomerCreateRequestBuilder) Phone(phone string) *CustomerCreateRequestBuilder {
	b.phone = &phone
	return b
}

// Metadata sets the metadata
func (b *CustomerCreateRequestBuilder) Metadata(metadata map[string]any) *CustomerCreateRequestBuilder {
	b.metadata = metadata
	return b
}

// Build creates the CustomerCreateRequest
func (b *CustomerCreateRequestBuilder) Build() *CustomerCreateRequest {
	return &CustomerCreateRequest{
		Email:     b.email,
		FirstName: b.firstName,
		LastName:  b.lastName,
		Phone:     b.phone,
		Metadata:  b.metadata,
	}
}

// Create creates a new customer with the provided builder
func (c *Client) Create(ctx context.Context, builder *CustomerCreateRequestBuilder) (*types.Response[types.Customer], error) {
	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()

	return net.Post[CustomerCreateRequest, types.Customer](
		ctx,
		c.client,
		c.secret,
		customerBasePath,
		req,
		c.baseURL,
	)
}
