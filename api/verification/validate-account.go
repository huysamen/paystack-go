package verification

import (
	"context"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AccountValidateRequest represents the request to validate an account
type AccountValidateRequest struct {
	AccountName    string  `json:"account_name"`              // Required: customer's name
	AccountNumber  string  `json:"account_number"`            // Required: account number
	AccountType    string  `json:"account_type"`              // Required: personal or business
	BankCode       string  `json:"bank_code"`                 // Required: bank code
	CountryCode    string  `json:"country_code"`              // Required: two digit ISO code
	DocumentType   string  `json:"document_type"`             // Required: identity document type
	DocumentNumber *string `json:"document_number,omitempty"` // Optional: identity document number
}

// AccountValidateRequestBuilder provides a fluent interface for building AccountValidateRequest
type AccountValidateRequestBuilder struct {
	req *AccountValidateRequest
}

// NewAccountValidateRequest creates a new builder for AccountValidateRequest
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

// DocumentNumber sets the optional document number
func (b *AccountValidateRequestBuilder) DocumentNumber(documentNumber string) *AccountValidateRequestBuilder {
	b.req.DocumentNumber = &documentNumber

	return b
}

// Build returns the constructed AccountValidateRequest
func (b *AccountValidateRequestBuilder) Build() *AccountValidateRequest {
	return b.req
}

// AccountValidateResponse represents the response from validating an account
type AccountValidateResponse = types.Response[types.AccountValidation]

// ValidateAccount validates an account using additional verification data
func (c *Client) ValidateAccount(ctx context.Context, builder *AccountValidateRequestBuilder) (*types.Response[types.AccountValidation], error) {
	return net.Post[AccountValidateRequest, types.AccountValidation](ctx, c.Client, c.Secret, accountValidateBasePath, builder.Build(), c.BaseURL)
}
