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

type CustomerValidateResponse struct {
	Message string `json:"message"`
}

// Builder for CustomerValidateRequest
type CustomerValidateRequestBuilder struct {
	firstName     string
	lastName      string
	type_         string
	value         string
	country       string
	bvn           string
	bankCode      string
	accountNumber string
	middleName    *string
}

// NewValidateCustomerRequest creates a new builder for customer validation
func NewValidateCustomerRequest(firstName, lastName, type_, value, country, bvn string) *CustomerValidateRequestBuilder {
	return &CustomerValidateRequestBuilder{
		firstName: firstName,
		lastName:  lastName,
		type_:     type_,
		value:     value,
		country:   country,
		bvn:       bvn,
	}
}

// BankCode sets the bank code (required if type is bank_account)
func (b *CustomerValidateRequestBuilder) BankCode(bankCode string) *CustomerValidateRequestBuilder {
	b.bankCode = bankCode
	return b
}

// AccountNumber sets the account number (required if type is bank_account)
func (b *CustomerValidateRequestBuilder) AccountNumber(accountNumber string) *CustomerValidateRequestBuilder {
	b.accountNumber = accountNumber
	return b
}

// MiddleName sets the middle name
func (b *CustomerValidateRequestBuilder) MiddleName(middleName string) *CustomerValidateRequestBuilder {
	b.middleName = &middleName
	return b
}

// Build creates the CustomerValidateRequest
func (b *CustomerValidateRequestBuilder) Build() *CustomerValidateRequest {
	return &CustomerValidateRequest{
		FirstName:     b.firstName,
		LastName:      b.lastName,
		Type:          b.type_,
		Value:         b.value,
		Country:       b.country,
		BVN:           b.bvn,
		BankCode:      b.bankCode,
		AccountNumber: b.accountNumber,
		MiddleName:    b.middleName,
	}
}

// Validate validates a customer with the provided builder
func (c *Client) Validate(ctx context.Context, customerCode string, builder *CustomerValidateRequestBuilder) (*types.Response[CustomerValidateResponse], error) {
	if customerCode == "" {
		return nil, fmt.Errorf("customer code is required")
	}

	if builder == nil {
		return nil, ErrBuilderRequired
	}

	req := builder.Build()
	path := fmt.Sprintf("%s/%s/identification", customerBasePath, customerCode)

	return net.Post[CustomerValidateRequest, CustomerValidateResponse](
		ctx,
		c.client,
		c.secret,
		path,
		req,
		c.baseURL,
	)
}
