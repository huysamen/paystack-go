package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/enums"
	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type partialDebitRequest struct {
	AuthorizationCode string         `json:"authorization_code"`
	Currency          enums.Currency `json:"currency"`
	Amount            int            `json:"amount"`
	Email             string         `json:"email"`

	Reference string `json:"reference,omitempty"`
	AtLeast   string `json:"at_least,omitempty"`
}

type PartialDebitRequestBuilder struct {
	request partialDebitRequest
}

func NewPartialDebitRequestBuilder() *PartialDebitRequestBuilder {
	return &PartialDebitRequestBuilder{}
}

func (b *PartialDebitRequestBuilder) AuthorizationCode(authorizationCode string) *PartialDebitRequestBuilder {
	b.request.AuthorizationCode = authorizationCode
	return b
}

func (b *PartialDebitRequestBuilder) Currency(currency enums.Currency) *PartialDebitRequestBuilder {
	b.request.Currency = currency
	return b
}

func (b *PartialDebitRequestBuilder) Amount(amount int) *PartialDebitRequestBuilder {
	b.request.Amount = amount
	return b
}

func (b *PartialDebitRequestBuilder) Email(email string) *PartialDebitRequestBuilder {
	b.request.Email = email
	return b
}

func (b *PartialDebitRequestBuilder) Reference(reference string) *PartialDebitRequestBuilder {
	b.request.Reference = reference
	return b
}

func (b *PartialDebitRequestBuilder) AtLeast(atLeast string) *PartialDebitRequestBuilder {
	b.request.AtLeast = atLeast
	return b
}

func (b *PartialDebitRequestBuilder) Build() *partialDebitRequest {
	return &b.request
}

type PartialDebitResponseData = types.Transaction
type PartialDebitResponse = types.Response[PartialDebitResponseData]

func (c *Client) PartialDebit(ctx context.Context, builder PartialDebitRequestBuilder) (*PartialDebitResponse, error) {
	return net.Post[partialDebitRequest, PartialDebitResponseData](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionPartialDebitPath), builder.Build(), c.BaseURL)
}
