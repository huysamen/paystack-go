package verification

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type validateAccountRequest struct {
	AccountName    string  `json:"account_name"`              // Required: customer's name
	AccountNumber  string  `json:"account_number"`            // Required: account number
	AccountType    string  `json:"account_type"`              // Required: personal or business
	BankCode       string  `json:"bank_code"`                 // Required: bank code
	CountryCode    string  `json:"country_code"`              // Required: two digit ISO code
	DocumentType   string  `json:"document_type"`             // Required: identity document type
	DocumentNumber *string `json:"document_number,omitempty"` // Optional: identity document number
}

type ValidateAccountRequestBuilder struct {
	req *validateAccountRequest
}

func NewValidateAccountRequestBuilder(accountName, accountNumber, accountType, bankCode, countryCode, documentType string) *ValidateAccountRequestBuilder {
	return &ValidateAccountRequestBuilder{
		req: &validateAccountRequest{
			AccountName:   accountName,
			AccountNumber: accountNumber,
			AccountType:   accountType,
			BankCode:      bankCode,
			CountryCode:   countryCode,
			DocumentType:  documentType,
		},
	}
}

func (b *ValidateAccountRequestBuilder) DocumentNumber(documentNumber string) *ValidateAccountRequestBuilder {
	b.req.DocumentNumber = &documentNumber

	return b
}

func (b *ValidateAccountRequestBuilder) Build() *validateAccountRequest {
	return b.req
}

type ValidateAccountResponseData = types.AccountValidation
type ValidateAccountResponse = types.Response[ValidateAccountResponseData]

func (c *Client) ValidateAccount(ctx context.Context, builder ValidateAccountRequestBuilder) (*ValidateAccountResponse, error) {
	return net.Post[validateAccountRequest, ValidateAccountResponseData](ctx, c.Client, c.Secret, accountValidateBasePath, builder.Build(), c.BaseURL)
}
