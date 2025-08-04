package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

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

type CustomerValidateRequestBuilder struct {
	req *CustomerValidateRequest
}

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

func (b *CustomerValidateRequestBuilder) BankCode(bankCode string) *CustomerValidateRequestBuilder {
	b.req.BankCode = bankCode

	return b
}

func (b *CustomerValidateRequestBuilder) AccountNumber(accountNumber string) *CustomerValidateRequestBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

func (b *CustomerValidateRequestBuilder) MiddleName(middleName string) *CustomerValidateRequestBuilder {
	b.req.MiddleName = &middleName

	return b
}

func (b *CustomerValidateRequestBuilder) Build() *CustomerValidateRequest {
	return b.req
}

type CustomerValidateResponse = types.Response[any]

func (c *Client) Validate(ctx context.Context, customerCode string, builder *CustomerValidateRequestBuilder) (*CustomerValidateResponse, error) {
	path := fmt.Sprintf("%s/%s/identification", basePath, customerCode)
	return net.Post[CustomerValidateRequest, any](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
