package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// CreateDedicatedVirtualAccountRequest represents the request to create a dedicated virtual account
type CreateDedicatedVirtualAccountRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type CreateDedicatedVirtualAccountResponse = types.Response[DedicatedVirtualAccount]

// CreateDedicatedVirtualAccountBuilder builds requests for creating dedicated virtual accounts
type CreateDedicatedVirtualAccountBuilder struct {
	request *CreateDedicatedVirtualAccountRequest
}

// NewCreateDedicatedVirtualAccountBuilder creates a new builder for creating dedicated virtual accounts
func NewCreateDedicatedVirtualAccountBuilder() *CreateDedicatedVirtualAccountBuilder {
	return &CreateDedicatedVirtualAccountBuilder{
		request: &CreateDedicatedVirtualAccountRequest{},
	}
}

// Customer sets the customer ID for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) Customer(customer string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Customer = customer
	return b
}

// PreferredBank sets the preferred bank for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) PreferredBank(preferredBank string) *CreateDedicatedVirtualAccountBuilder {
	b.request.PreferredBank = preferredBank
	return b
}

// Subaccount sets the subaccount for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) Subaccount(subaccount string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Subaccount = subaccount
	return b
}

// SplitCode sets the split code for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) SplitCode(splitCode string) *CreateDedicatedVirtualAccountBuilder {
	b.request.SplitCode = splitCode
	return b
}

// FirstName sets the first name for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) FirstName(firstName string) *CreateDedicatedVirtualAccountBuilder {
	b.request.FirstName = firstName
	return b
}

// LastName sets the last name for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) LastName(lastName string) *CreateDedicatedVirtualAccountBuilder {
	b.request.LastName = lastName
	return b
}

// Phone sets the phone for the dedicated virtual account
func (b *CreateDedicatedVirtualAccountBuilder) Phone(phone string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Phone = phone
	return b
}

// Build returns the built request
func (b *CreateDedicatedVirtualAccountBuilder) Build() *CreateDedicatedVirtualAccountRequest {
	return b.request
}

// Create creates a dedicated virtual account for an existing customer
func (c *Client) Create(ctx context.Context, builder *CreateDedicatedVirtualAccountBuilder) (*CreateDedicatedVirtualAccountResponse, error) {
	return net.Post[CreateDedicatedVirtualAccountRequest, DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
