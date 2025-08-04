package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type assignRequest struct {
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

type AssignRequestBuilder struct {
	req *assignRequest
}

func NewAssignRequestBuilder() *AssignRequestBuilder {
	return &AssignRequestBuilder{
		req: &assignRequest{},
	}
}

func (b *AssignRequestBuilder) Email(email string) *AssignRequestBuilder {
	b.req.Email = email

	return b
}

func (b *AssignRequestBuilder) FirstName(firstName string) *AssignRequestBuilder {
	b.req.FirstName = firstName

	return b
}

func (b *AssignRequestBuilder) LastName(lastName string) *AssignRequestBuilder {
	b.req.LastName = lastName

	return b
}

func (b *AssignRequestBuilder) Phone(phone string) *AssignRequestBuilder {
	b.req.Phone = phone

	return b
}

func (b *AssignRequestBuilder) PreferredBank(preferredBank string) *AssignRequestBuilder {
	b.req.PreferredBank = preferredBank

	return b
}

func (b *AssignRequestBuilder) Country(country string) *AssignRequestBuilder {
	b.req.Country = country

	return b
}

func (b *AssignRequestBuilder) AccountNumber(accountNumber string) *AssignRequestBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

func (b *AssignRequestBuilder) BVN(bvn string) *AssignRequestBuilder {
	b.req.BVN = bvn

	return b
}

func (b *AssignRequestBuilder) BankCode(bankCode string) *AssignRequestBuilder {
	b.req.BankCode = bankCode

	return b
}

func (b *AssignRequestBuilder) Subaccount(subaccount string) *AssignRequestBuilder {
	b.req.Subaccount = subaccount

	return b
}

func (b *AssignRequestBuilder) SplitCode(splitCode string) *AssignRequestBuilder {
	b.req.SplitCode = splitCode

	return b
}

func (b *AssignRequestBuilder) MiddleName(middleName string) *AssignRequestBuilder {
	b.req.MiddleName = middleName

	return b
}

func (b *AssignRequestBuilder) Build() *assignRequest {
	return b.req
}

type AssignDedicatedVirtualAccountResponseData = any
type AssignDedicatedVirtualAccountResponse = types.Response[AssignDedicatedVirtualAccountResponseData]

func (c *Client) Assign(ctx context.Context, builder AssignRequestBuilder) (*AssignDedicatedVirtualAccountResponse, error) {
	return net.Post[assignRequest, AssignDedicatedVirtualAccountResponseData](ctx, c.Client, c.Secret, basePath+"/assign", builder.Build(), c.BaseURL)
}
