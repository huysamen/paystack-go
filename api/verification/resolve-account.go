package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

// AccountResolveRequest represents the request to resolve an account number
type AccountResolveRequest struct {
	AccountNumber string `json:"account_number"` // Required: account number
	BankCode      string `json:"bank_code"`      // Required: bank code
}

// AccountResolveRequestBuilder builds an AccountResolveRequest
type AccountResolveRequestBuilder struct {
	request AccountResolveRequest
}

// NewAccountResolveRequestBuilder creates a new builder
func NewAccountResolveRequestBuilder() *AccountResolveRequestBuilder {
	return &AccountResolveRequestBuilder{}
}

// AccountNumber sets the account number
func (b *AccountResolveRequestBuilder) AccountNumber(accountNumber string) *AccountResolveRequestBuilder {
	b.request.AccountNumber = accountNumber
	return b
}

// BankCode sets the bank code
func (b *AccountResolveRequestBuilder) BankCode(bankCode string) *AccountResolveRequestBuilder {
	b.request.BankCode = bankCode
	return b
}

// Build returns the built AccountResolveRequest
func (b *AccountResolveRequestBuilder) Build() *AccountResolveRequest {
	return &b.request
}

// AccountResolveResponse represents the response from resolving an account
type AccountResolveResponse = types.Response[types.AccountResolution]

// ResolveAccount resolves a bank account number to get account details
func (c *Client) ResolveAccount(ctx context.Context, builder *AccountResolveRequestBuilder) (*types.Response[types.AccountResolution], error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s?account_number=%s&bank_code=%s", accountResolveBasePath, req.AccountNumber, req.BankCode)
	return net.Get[types.AccountResolution](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
