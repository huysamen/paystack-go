package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateDedicatedVirtualAccountRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type CreateDedicatedVirtualAccountBuilder struct {
	request *CreateDedicatedVirtualAccountRequest
}

func NewCreateDedicatedVirtualAccountBuilder() *CreateDedicatedVirtualAccountBuilder {
	return &CreateDedicatedVirtualAccountBuilder{
		request: &CreateDedicatedVirtualAccountRequest{},
	}
}

func (b *CreateDedicatedVirtualAccountBuilder) Customer(customer string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Customer = customer

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) PreferredBank(preferredBank string) *CreateDedicatedVirtualAccountBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) Subaccount(subaccount string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) SplitCode(splitCode string) *CreateDedicatedVirtualAccountBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) FirstName(firstName string) *CreateDedicatedVirtualAccountBuilder {
	b.request.FirstName = firstName

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) LastName(lastName string) *CreateDedicatedVirtualAccountBuilder {
	b.request.LastName = lastName

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) Phone(phone string) *CreateDedicatedVirtualAccountBuilder {
	b.request.Phone = phone

	return b
}

func (b *CreateDedicatedVirtualAccountBuilder) Build() *CreateDedicatedVirtualAccountRequest {
	return b.request
}

type CreateDedicatedVirtualAccountResponse = types.Response[types.DedicatedVirtualAccount]

func (c *Client) Create(ctx context.Context, builder *CreateDedicatedVirtualAccountBuilder) (*CreateDedicatedVirtualAccountResponse, error) {
	return net.Post[CreateDedicatedVirtualAccountRequest, types.DedicatedVirtualAccount](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
