package dedicatedvirtualaccount

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AssignDedicatedVirtualAccountRequest represents the request to assign a dedicated virtual account
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

// AssignDedicatedVirtualAccountBuilder builds requests for assigning dedicated virtual accounts
type AssignDedicatedVirtualAccountBuilder struct {
	req *AssignDedicatedVirtualAccountRequest
}

// NewAssignDedicatedVirtualAccountBuilder creates a new builder for assigning dedicated virtual accounts
func NewAssignDedicatedVirtualAccountBuilder() *AssignDedicatedVirtualAccountBuilder {
	return &AssignDedicatedVirtualAccountBuilder{
		req: &AssignDedicatedVirtualAccountRequest{},
	}
}

// Email sets the email for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) Email(email string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Email = email

	return b
}

// FirstName sets the first name for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) FirstName(firstName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.FirstName = firstName

	return b
}

// LastName sets the last name for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) LastName(lastName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.LastName = lastName

	return b
}

// Phone sets the phone for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) Phone(phone string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Phone = phone

	return b
}

// PreferredBank sets the preferred bank for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) PreferredBank(preferredBank string) *AssignDedicatedVirtualAccountBuilder {
	b.req.PreferredBank = preferredBank

	return b
}

// Country sets the country for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) Country(country string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Country = country

	return b
}

// AccountNumber sets the account number for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) AccountNumber(accountNumber string) *AssignDedicatedVirtualAccountBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

// BVN sets the BVN for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) BVN(bvn string) *AssignDedicatedVirtualAccountBuilder {
	b.req.BVN = bvn

	return b
}

// BankCode sets the bank code for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) BankCode(bankCode string) *AssignDedicatedVirtualAccountBuilder {
	b.req.BankCode = bankCode

	return b
}

// Subaccount sets the subaccount for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) Subaccount(subaccount string) *AssignDedicatedVirtualAccountBuilder {
	b.req.Subaccount = subaccount

	return b
}

// SplitCode sets the split code for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) SplitCode(splitCode string) *AssignDedicatedVirtualAccountBuilder {
	b.req.SplitCode = splitCode

	return b
}

// MiddleName sets the middle name for assigning the dedicated virtual account
func (b *AssignDedicatedVirtualAccountBuilder) MiddleName(middleName string) *AssignDedicatedVirtualAccountBuilder {
	b.req.MiddleName = middleName

	return b
}

// Build returns the built request
func (b *AssignDedicatedVirtualAccountBuilder) Build() *AssignDedicatedVirtualAccountRequest {
	return b.req
}

// AssignDedicatedVirtualAccountResponse represents the response for assigning a dedicated virtual account
type AssignDedicatedVirtualAccountResponse = types.Response[any]

// Assign creates a customer, validates the customer, and assigns a dedicated virtual account
func (c *Client) Assign(ctx context.Context, builder *AssignDedicatedVirtualAccountBuilder) (*AssignDedicatedVirtualAccountResponse, error) {
	return net.Post[AssignDedicatedVirtualAccountRequest, any](
		ctx, c.Client, c.Secret, basePath+"/assign", builder.Build(), c.BaseURL,
	)
}
