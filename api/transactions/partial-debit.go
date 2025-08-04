package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionPartialDebitRequest struct {
	AuthorizationCode string         `json:"authorization_code"`
	Currency          types.Currency `json:"currency"`
	Amount            int            `json:"amount"`
	Email             string         `json:"email"`

	Reference string `json:"reference,omitempty"`
	AtLeast   string `json:"at_least,omitempty"`
}

type TransactionPartialDebitRequestBuilder struct {
	request TransactionPartialDebitRequest
}

func NewTransactionPartialDebitRequestBuilder() *TransactionPartialDebitRequestBuilder {
	return &TransactionPartialDebitRequestBuilder{}
}

func (b *TransactionPartialDebitRequestBuilder) AuthorizationCode(authorizationCode string) *TransactionPartialDebitRequestBuilder {
	b.request.AuthorizationCode = authorizationCode
	return b
}

func (b *TransactionPartialDebitRequestBuilder) Currency(currency types.Currency) *TransactionPartialDebitRequestBuilder {
	b.request.Currency = currency
	return b
}

func (b *TransactionPartialDebitRequestBuilder) Amount(amount int) *TransactionPartialDebitRequestBuilder {
	b.request.Amount = amount
	return b
}

func (b *TransactionPartialDebitRequestBuilder) Email(email string) *TransactionPartialDebitRequestBuilder {
	b.request.Email = email
	return b
}

func (b *TransactionPartialDebitRequestBuilder) Reference(reference string) *TransactionPartialDebitRequestBuilder {
	b.request.Reference = reference
	return b
}

func (b *TransactionPartialDebitRequestBuilder) AtLeast(atLeast string) *TransactionPartialDebitRequestBuilder {
	b.request.AtLeast = atLeast
	return b
}

func (b *TransactionPartialDebitRequestBuilder) Build() *TransactionPartialDebitRequest {
	return &b.request
}

type PartialDebitResponse = types.Response[types.Transaction]

func (c *Client) PartialDebit(ctx context.Context, builder *TransactionPartialDebitRequestBuilder) (*PartialDebitResponse, error) {
	return net.Post[TransactionPartialDebitRequest, types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionPartialDebitPath), builder.Build(), c.BaseURL)
}
