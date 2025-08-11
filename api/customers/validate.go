package customers

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type validateRequest struct {
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

type ValidateRequestBuilder struct {
	req *validateRequest
}

func NewValidateRequestBuilder(firstName, lastName, type_, value, country, bvn string) *ValidateRequestBuilder {
	return &ValidateRequestBuilder{
		req: &validateRequest{
			FirstName: firstName,
			LastName:  lastName,
			Type:      type_,
			Value:     value,
			Country:   country,
			BVN:       bvn,
		},
	}
}

func (b *ValidateRequestBuilder) BankCode(bankCode string) *ValidateRequestBuilder {
	b.req.BankCode = bankCode

	return b
}

func (b *ValidateRequestBuilder) AccountNumber(accountNumber string) *ValidateRequestBuilder {
	b.req.AccountNumber = accountNumber

	return b
}

func (b *ValidateRequestBuilder) MiddleName(middleName string) *ValidateRequestBuilder {
	b.req.MiddleName = &middleName

	return b
}

func (b *ValidateRequestBuilder) Build() *validateRequest {
	return b.req
}

type ValidateResponseData = any
type CustomerValidateResponse = types.Response[ValidateResponseData]

func (c *Client) Validate(ctx context.Context, customerCode string, builder ValidateRequestBuilder) (*CustomerValidateResponse, error) {
	path := fmt.Sprintf("%s/%s/identification", basePath, customerCode)

	return net.Post[validateRequest, ValidateResponseData](ctx, c.Client, c.Secret, path, builder.Build(), c.BaseURL)
}
