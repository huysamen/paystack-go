package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type CreateRequest struct {
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type CreateRequestBuilder struct {
	request *CreateRequest
}

func NewCreateRequestBuilder() *CreateRequestBuilder {
	return &CreateRequestBuilder{
		request: &CreateRequest{},
	}
}

func (b *CreateRequestBuilder) Customer(customer string) *CreateRequestBuilder {
	b.request.Customer = customer

	return b
}

func (b *CreateRequestBuilder) PreferredBank(preferredBank string) *CreateRequestBuilder {
	b.request.PreferredBank = preferredBank

	return b
}

func (b *CreateRequestBuilder) Subaccount(subaccount string) *CreateRequestBuilder {
	b.request.Subaccount = subaccount

	return b
}

func (b *CreateRequestBuilder) SplitCode(splitCode string) *CreateRequestBuilder {
	b.request.SplitCode = splitCode

	return b
}

func (b *CreateRequestBuilder) FirstName(firstName string) *CreateRequestBuilder {
	b.request.FirstName = firstName

	return b
}

func (b *CreateRequestBuilder) LastName(lastName string) *CreateRequestBuilder {
	b.request.LastName = lastName

	return b
}

func (b *CreateRequestBuilder) Phone(phone string) *CreateRequestBuilder {
	b.request.Phone = phone

	return b
}

func (b *CreateRequestBuilder) Build() *CreateRequest {
	return b.request
}

type CreateResponseData = types.DedicatedVirtualAccount
type CreateResponse = types.Response[CreateResponseData]

func (c *Client) Create(ctx context.Context, builder CreateRequestBuilder) (*CreateResponse, error) {
	return net.Post[CreateRequest, CreateResponseData](ctx, c.Client, c.Secret, basePath, builder.Build(), c.BaseURL)
}
