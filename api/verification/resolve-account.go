package verification

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type resolveAccountRequest struct {
	AccountNumber string `json:"account_number"` // Required: account number
	BankCode      string `json:"bank_code"`      // Required: bank code
}

type ResolveAccountRequestBuilder struct {
	request resolveAccountRequest
}

func NewResolveAccountRequestBuilder() *ResolveAccountRequestBuilder {
	return &ResolveAccountRequestBuilder{}
}

func (b *ResolveAccountRequestBuilder) AccountNumber(accountNumber string) *ResolveAccountRequestBuilder {
	b.request.AccountNumber = accountNumber

	return b
}

func (b *ResolveAccountRequestBuilder) BankCode(bankCode string) *ResolveAccountRequestBuilder {
	b.request.BankCode = bankCode

	return b
}

func (b *ResolveAccountRequestBuilder) Build() *resolveAccountRequest {
	return &b.request
}

type ResolveAccountResponseData = types.AccountResolution
type ResolveAccountResponse = types.Response[ResolveAccountResponseData]

func (c *Client) ResolveAccount(ctx context.Context, builder ResolveAccountRequestBuilder) (*ResolveAccountResponse, error) {
	req := builder.Build()
	endpoint := fmt.Sprintf("%s?account_number=%s&bank_code=%s", accountResolveBasePath, req.AccountNumber, req.BankCode)

	return net.Get[ResolveAccountResponseData](ctx, c.Client, c.Secret, endpoint, "", c.BaseURL)
}
