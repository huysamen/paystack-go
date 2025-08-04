package transactions

import (
	"context"
	"fmt"

	"github.com/huysamen/paystack-go/net"
	"github.com/huysamen/paystack-go/types"
)

type TransactionPartialDebitRequest struct {
	// Required
	AuthorizationCode string         `json:"authorization_code"`
	Currency          types.Currency `json:"currency"`
	Amount            int            `json:"amount"`
	Email             string         `json:"email"`

	// Optional
	Reference string `json:"reference,omitempty"`
	AtLeast   string `json:"at_least,omitempty"`
}

// TransactionPartialDebitRequestBuilder builds a TransactionPartialDebitRequest
type TransactionPartialDebitRequestBuilder struct {
	request TransactionPartialDebitRequest
}

// NewTransactionPartialDebitRequestBuilder creates a new builder
func NewTransactionPartialDebitRequestBuilder() *TransactionPartialDebitRequestBuilder {
	return &TransactionPartialDebitRequestBuilder{}
}

// AuthorizationCode sets the authorization code for the transaction
func (b *TransactionPartialDebitRequestBuilder) AuthorizationCode(authorizationCode string) *TransactionPartialDebitRequestBuilder {
	b.request.AuthorizationCode = authorizationCode
	return b
}

// Currency sets the currency for the transaction
func (b *TransactionPartialDebitRequestBuilder) Currency(currency types.Currency) *TransactionPartialDebitRequestBuilder {
	b.request.Currency = currency
	return b
}

// Amount sets the amount for the transaction
func (b *TransactionPartialDebitRequestBuilder) Amount(amount int) *TransactionPartialDebitRequestBuilder {
	b.request.Amount = amount
	return b
}

// Email sets the email for the transaction
func (b *TransactionPartialDebitRequestBuilder) Email(email string) *TransactionPartialDebitRequestBuilder {
	b.request.Email = email
	return b
}

// Reference sets the reference for the transaction
func (b *TransactionPartialDebitRequestBuilder) Reference(reference string) *TransactionPartialDebitRequestBuilder {
	b.request.Reference = reference
	return b
}

// AtLeast sets the at least amount for the transaction
func (b *TransactionPartialDebitRequestBuilder) AtLeast(atLeast string) *TransactionPartialDebitRequestBuilder {
	b.request.AtLeast = atLeast
	return b
}

// Build returns the built TransactionPartialDebitRequest
func (b *TransactionPartialDebitRequestBuilder) Build() *TransactionPartialDebitRequest {
	return &b.request
}

// Response type alias
type PartialDebitResponse = types.Response[types.Transaction]

func (c *Client) PartialDebit(ctx context.Context, builder *TransactionPartialDebitRequestBuilder) (*PartialDebitResponse, error) {
	return net.Post[TransactionPartialDebitRequest, types.Transaction](ctx, c.Client, c.Secret, fmt.Sprintf("%s%s", basePath, transactionPartialDebitPath), builder.Build(), c.BaseURL)
}
