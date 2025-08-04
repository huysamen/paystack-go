package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AssignDedicatedVirtualAccountRequest struct {
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Phone         string `json:"phone"`
	PreferredBank string `json:"preferred_bank"`
	Country       string `json:"country"`
	AccountNumber string `json:"account_number,omitempty"`
	BVN           string `json:"bvn,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
}

type AssignDedicatedVirtualAccountBuilder struct {
	req *AssignDedicatedVirtualAccountRequest
}

func NewAssignDedicatedVirtualAccountBuilder() *AssignDedicatedVirtualAccountBuilder {
	return &AssignDedicatedVirtualAccountBuilder{
		req: &AssignDedicatedVirtualAccountRequest{},
	}
}

func (b *AssignDedicatedVirtualAccountBuilder) Email(email string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Email = email

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) FirstName(firstName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.FirstName = firstName

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) LastName(lastName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.LastName = lastName

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) Phone(phone string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Phone = phone

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) PreferredBank(preferredBank string) *AssignDedicatedVirtualAccountBuilder {
	b.req.PreferredBank = preferredBank

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) Country(country string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Country = country

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) AccountNumber(accountNumber string) *AssignDedicatedVirtualAccountBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) BVN(bvn string) *AssignDedicatedVirtualAccountBuilder {
	b.req.BVN = bvn

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) BankCode(bankCode string) *AssignDedicatedVirtualAccountBuilder {
	b.req.BankCode = bankCode

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) Subaccount(subaccount string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Subaccount = subaccount

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) SplitCode(splitCode string) *AssignDedicatedVirtualAccountBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) MiddleName(middleName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.MiddleName = middleName

	return b
}

func (b *AssignDedicatedVirtualAccountBuilder) Build() *AssignDedicatedVirtualAccountRequest {
	return b.req
}

type AssignDedicatedVirtualAccountResponse = types.Response[any]

func (c *Client) Assign(ctx context.Context, builder *AssignDedicatedVirtualAccountBuilder) (*AssignDedicatedVirtualAccountResponse, error) {
	return net.Post[AssignDedicatedVirtualAccountRequest, any](
		ctx, c.Client, c.Secret, basePath+"/assign", builder.Build(), c.BaseURL,
	)
}
