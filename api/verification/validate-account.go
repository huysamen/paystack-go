package verification

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AccountValidateRequest struct {
	AccountName    string  `json:"account_name"`              // Required: customer's name
	AccountNumber  string  `json:"account_number"`            // Required: account number
	AccountType    string  `json:"account_type"`              // Required: personal or business
	BankCode       string  `json:"bank_code"`                 // Required: bank code
	CountryCode    string  `json:"country_code"`              // Required: two digit ISO code
	DocumentType   string  `json:"document_type"`             // Required: identity document type
	DocumentNumber *string `json:"document_number,omitempty"` // Optional: identity document number
}

type AccountValidateRequestBuilder struct {
	req *AccountValidateRequest
}

func NewAccountValidateRequest(accountName, accountNumber, accountType, bankCode, countryCode, documentType string) *AccountValidateRequestBuilder {
	return &AccountValidateRequestBuilder{
		req: &AccountValidateRequest{
			AccountName:   accountName,
			AccountNumber: accountNumber,
			AccountType:   accountType,
			BankCode:      bankCode,
			CountryCode:   countryCode,
			DocumentType:  documentType,
		},
	}
}

func (b *AccountValidateRequestBuilder) DocumentNumber(documentNumber string) *AccountValidateRequestBuilder {
	b.req.DocumentNumber = &documentNumber

	return b
}

func (b *AccountValidateRequestBuilder) Build() *AccountValidateRequest {
	return b.req
}

type AccountValidateResponse = types.Response[types.AccountValidation]

func (c *Client) ValidateAccount(ctx context.Context, builder *AccountValidateRequestBuilder) (*types.Response[types.AccountValidation], error) {
	return net.Post[AccountValidateRequest, types.AccountValidation](ctx, c.Client, c.Secret, accountValidateBasePath, builder.Build(), c.BaseURL)
}
