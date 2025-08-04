package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// Request and Response types
type CustomerValidateRequest struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Type          string  `json:"type"` // Only "bank_account" is supported
	Value         string  `json:"value"`
	Country       string  `json:"country"`
	BVN           string  `json:"bvn"`
	BankCode      string  `json:"bank_code"`      // Required if type is bank_account
	AccountNumber string  `json:"account_number"` // Required if type is bank_account
	MiddleName    *string `json:"middle_name,omitempty"`
}

// Builder for CustomerValidateRequest
type CustomerValidateRequestBuilder struct {
	req *CustomerValidateRequest
}

// NewValidateCustomerRequest creates a new builder for customer validation
func NewValidateCustomerRequest(firstName, lastName, type_, value, country, bvn string) *CustomerValidateRequestBuilder {
	return &CustomerValidateRequestBuilder{
		req: &CustomerValidateRequest{
			FirstName: firstName,
			LastName:  lastName,
			Type:      type_,
			Value:     value,
			Country:   country,
			BVN:       bvn,
		},
	}
}

// BankCode sets the bank code (required if type is bank_account)
func (b *CustomerValidateRequestBuilder) BankCode(bankCode string) *CustomerValidateRequestBuilder {
	b.req.BankCode = bankCode

	return b
}

// AccountNumber sets the account number (required if type is bank_account)
func (b *CustomerValidateRequestBuilder) AccountNumber(accountNumber string) *CustomerValidateRequestBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

// MiddleName sets the middle name
func (b *CustomerValidateRequestBuilder) MiddleName(middleName string) *CustomerValidateRequestBuilder {
	b.req.MiddleName = &middleName

	return b
}

// Build creates the CustomerValidateRequest
func (b *CustomerValidateRequestBuilder) Build() *CustomerValidateRequest {
	return b.req
}

// CustomerValidateResponse represents the response type for customer validation
type CustomerValidateResponse = types.Response[any]

// Validate validates a customer with the provided builder
func (c *Client) Validate(ctx context.Context, customerCode string, builder *CustomerValidateRequestBuilder) (*CustomerValidateResponse, error) {
	path := fmt.Sprintf("%s/%s/identification", basePath, customerCode)
	return net.Post[CustomerValidateRequest, any](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
