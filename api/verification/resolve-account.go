package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type AccountResolveRequest struct {
	AccountNumber string `json:"account_number"` // Required: account number
	BankCode      string `json:"bank_code"`      // Required: bank code
}

type AccountResolveRequestBuilder struct {
	request AccountResolveRequest
}

func NewAccountResolveRequestBuilder() *AccountResolveRequestBuilder {
	return &AccountResolveRequestBuilder{}
}

func (b *AccountResolveRequestBuilder) AccountNumber(accountNumber string) *AccountResolveRequestBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *AccountResolveRequestBuilder) BankCode(bankCode string) *AccountResolveRequestBuilder {
	b.request.BankCode = bankCode

	return b
}

func (b *AccountResolveRequestBuilder) Build() *AccountResolveRequest {
	return &b.request
}

type AccountResolveResponse = types.Response[types.AccountResolution]

func (c *Client) ResolveAccount(ctx context.Context, builder *AccountResolveRequestBuilder) (*AccountResolveResponse, error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s?account_number=%s&bank_code=%s", accountResolveBasePath, req.AccountNumber, req.BankCode)
	return net.Get[types.AccountResolution](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
